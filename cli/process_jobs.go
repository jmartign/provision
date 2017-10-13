package cli

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
	"time"

	"github.com/digitalrebar/provision/client/events"
	"github.com/digitalrebar/provision/client/jobs"
	"github.com/digitalrebar/provision/client/machines"
	models "github.com/digitalrebar/provision/genmodels"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

var exitOnFailure = false

func Log(uuid *strfmt.UUID, s string) error {
	buf := bytes.NewBufferString(s)
	_, err := session.Jobs.PutJobLog(jobs.NewPutJobLogParams().WithUUID(*uuid).WithBody(buf), basicAuth)
	if err != nil {
		fmt.Printf("Failed to log to job log, %s: %v\n", uuid.String(), err)
	}
	return err
}

func writeStringToFile(filename, content string) error {
	dir := path.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	fo, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(content))
	if err != nil {
		return err
	}
	return nil
}

func markJob(uuid, state string, ops ModOps) error {
	j := fmt.Sprintf("{\"State\": \"%s\"}", state)
	if _, err := Update(uuid, j, ops, false); err != nil {
		fmt.Printf("Error marking job, %s, as %s: %v, continuing\n", uuid, state, err)
		return err
	}
	return nil
}

func markMachineRunnable(uuid string, ops ModOps) error {
	if _, err := Update(uuid, `{"Runnable": true}`, ops, false); err != nil {
		fmt.Printf("Error marking machine as runnable: %v, continuing to wait for runnable...\n", err)
		return err
	}
	return nil
}

func setMachineCurrentTask(uuid string, ct int, ops ModOps) error {
	if _, err := Update(uuid, fmt.Sprintf("{\"CurrentTask\": %d}", ct), ops, false); err != nil {
		fmt.Printf("Error setting machine as currentTask: %v\n", err)
		return err
	}
	return nil
}

func setMachineStage(uuid, stage string, ops ModOps) error {
	j := fmt.Sprintf("{\"Stage\": \"%s\"}", stage)
	oldForce := force
	force = true
	if _, err := Update(uuid, j, ops, false); err != nil {
		fmt.Printf("Error setting stage on machine, %s, as %s: %v, continuing\n", uuid, stage, err)
		return err
	}
	force = oldForce
	return nil
}

func setMachineWorkflow(uuid, workflow string) error {
	return setMachineParam(uuid, "workflows/current", workflow)
}

func clearMachineWorkflow(uuid string) error {
	return setMachineParam(uuid, "workflows/current", nil)
}

type CommandRunner struct {
	name string
	uuid *strfmt.UUID

	cmd      *exec.Cmd
	stderr   io.ReadCloser
	stdout   io.ReadCloser
	stdin    io.WriteCloser
	finished chan bool
}

func (cr *CommandRunner) ReadLog() {
	// read command's stderr line by line - for logging
	in := bufio.NewScanner(cr.stderr)
	for in.Scan() {
		Log(cr.uuid, in.Text())
	}
	cr.finished <- true
}

func (cr *CommandRunner) ReadReply() {
	// read command's stdout line by line - for replies
	in := bufio.NewScanner(cr.stdout)
	for in.Scan() {
		Log(cr.uuid, in.Text())
	}
	cr.finished <- true
}

func (cr *CommandRunner) Run() (failed, incomplete, reboot bool) {
	// Start command running
	err := cr.cmd.Start()
	if err != nil {
		failed = true
		reboot = false
		s := fmt.Sprintf("Command %s failed to start: %v\n", cr.name, err)
		fmt.Printf(s)
		Log(cr.uuid, s)
		return
	}

	// Wait for readers to exit
	<-cr.finished
	<-cr.finished

	err = cr.cmd.Wait()
	if exiterr, ok := err.(*exec.ExitError); ok {
		// The program has exited with an exit code != 0

		// This works on both Unix and Windows. Although package
		// syscall is generally platform dependent, WaitStatus is
		// defined for both Unix and Windows and in both cases has
		// an ExitStatus() method with the same signature.
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			code := status.ExitStatus()
			switch code {
			case 0:
				failed = false
				reboot = false
				s := fmt.Sprintf("Command %s succeeded\n", cr.name)
				fmt.Printf(s)
				Log(cr.uuid, s)
			case 1:
				failed = false
				reboot = true
				s := fmt.Sprintf("Command %s succeeded (wants reboot)\n", cr.name)
				fmt.Printf(s)
				Log(cr.uuid, s)
			case 2:
				incomplete = true
				reboot = false
				s := fmt.Sprintf("Command %s incomplete\n", cr.name)
				fmt.Printf(s)
				Log(cr.uuid, s)
			case 3:
				incomplete = true
				reboot = true
				s := fmt.Sprintf("Command %s incomplete (wants reboot)\n", cr.name)
				fmt.Printf(s)
				Log(cr.uuid, s)
			default:
				failed = true
				reboot = false
				s := fmt.Sprintf("Command %s failed\n", cr.name)
				fmt.Printf(s)
				Log(cr.uuid, s)
			}
		}
	} else {
		if err != nil {
			failed = true
			reboot = false
			s := fmt.Sprintf("Command %s failed: %v\n", cr.name, err)
			fmt.Printf(s)
			Log(cr.uuid, s)
		} else {
			failed = false
			reboot = false
			s := fmt.Sprintf("Command %s succeeded\n", cr.name)
			fmt.Printf(s)
			Log(cr.uuid, s)
		}
	}

	// Remove script
	os.Remove(cr.cmd.Path)

	return
}

func NewCommandRunner(uuid *strfmt.UUID, name, content string) (*CommandRunner, error) {
	answer := &CommandRunner{name: name, uuid: uuid}

	// Make script file
	tmpFile, err := ioutil.TempFile(".", "script")
	if err != nil {
		return nil, err
	}
	if _, err := tmpFile.Write([]byte(content)); err != nil {
		return nil, err
	}
	path := "./" + tmpFile.Name()
	if err := tmpFile.Close(); err != nil {
		return nil, err
	}
	os.Chmod(path, 0700)

	answer.cmd = exec.Command(path)

	var err2 error
	answer.stderr, err2 = answer.cmd.StderrPipe()
	if err2 != nil {
		return nil, err2
	}
	answer.stdout, err2 = answer.cmd.StdoutPipe()
	if err2 != nil {
		return nil, err2
	}
	answer.stdin, err2 = answer.cmd.StdinPipe()
	if err2 != nil {
		return nil, err2
	}

	answer.finished = make(chan bool, 2)
	go answer.ReadLog()
	go answer.ReadReply()

	return answer, nil
}

func runContent(uuid *strfmt.UUID, action *models.JobAction) (failed, incomplete, reboot bool) {

	Log(uuid, fmt.Sprintf("Starting Content Execution for: %s\n", *action.Name))

	runner, err := NewCommandRunner(uuid, *action.Name, *action.Content)
	if err != nil {
		failed = true
		s := fmt.Sprintf("Creating command %s failed: %v\n", *action.Name, err)
		fmt.Printf(s)
		Log(uuid, s)
	} else {
		failed, incomplete, reboot = runner.Run()
	}
	return
}

func getMachineParams(uuid string) (map[string]interface{}, error) {
	// Get all aggregate parameters.
	as := "true"
	d, err := session.Machines.GetMachineParams(machines.NewGetMachineParamsParams().WithAggregate(&as).WithUUID(strfmt.UUID(uuid)), basicAuth)
	if err != nil {
		return nil, generateError(err, "Failed to fetch params: %v", uuid)
	}
	return d.Payload, nil
}

func getWorkflowsCurrent(params map[string]interface{}) (string, bool) {
	if wobj, ok := params["workflows/current"]; !ok {
		return "", false
	} else {
		s, ok := wobj.(string)
		return s, ok
	}
}
func getWorkflowsVersion(params map[string]interface{}) (string, bool) {
	if wobj, ok := params["workflows/version"]; !ok {
		return "", false
	} else {
		s, ok := wobj.(string)
		return s, ok
	}
}

func getWorkflowsWorkflow(params map[string]interface{}, workflow string) (map[string]string, bool) {
	if obj, ok := params["workflows/map"]; !ok {
		return nil, false
	} else {
		cmap, ok := obj.(map[string]interface{})
		if !ok {
			return nil, false
		}

		// Get map for current workflow
		wfobj, ok := cmap[workflow]
		if !ok {
			return nil, false
		}
		s, ok := wfobj.(map[string]string)
		return s, ok
	}
}

func getChangeStageMap(params map[string]interface{}) (map[string]string, bool) {
	if obj, ok := params["change-stage/map"]; !ok {
		return map[string]string{}, false
	} else {
		s, ok := obj.(map[string]string)
		return s, ok
	}
}

func findCurrentWorkflow(params map[string]interface{}, stage string) (string, bool) {
	if obj, ok := params["workflows/map"]; !ok {
		return nil, false
	} else {
		cmap, ok := obj.(map[string]interface{})
		if !ok {
			return nil, false
		}

		for _, wf := range cmap {
			firstStage, err := getWorkflowsWorkflowFirstStage(params, wf)
			if err != nil {
				continue
			}
			if firstStage == stage {
				return workflow, true
			}
		}
	}
	return nil, false
}

func getWorkflowsWorkflowFirstStage(params map[string]interface{}, workflow string) (string, error) {
	wfmap, err := getWorkflowsWorkflow(params, workflow)

	// We now have a map of Sx -> Sy links.  We need to find the root of this tree.

	// GREG: Fix this
	root := "GREG"

	return "", nil
}

func getNextStage(uuid, stage string, so *StageOps) (string, bool, error) {
	params, err := getMachineParams(uuid)
	if err != nil {
		return "", false, err
	}

	// Find next stage
	// - check for workflow/version
	// - if no version, check for old stage map
	// - else do version methods.
	nextStage := ""
	reboot := false

	version, ok := getWorkflowsVersion(params)
	if !ok {
		cmap, ok := getChangeStageMap(params)
		if !ok {
			return "", false, nil
		} else {
			if ns, ok := cmap[stage]; ok {
				pieces := strings.Split(ns, ":")
				nextStage = pieces[0]
				if len(pieces) > 1 && pieces[1] == "Reboot" {
					reboot = true
				}
			} else {
				nextStage = ""
			}
		}
	} else {
		if version != "1.0" {
			return "", false, generateError(err, "Unknown workflow version %s: %s", version, uuid)
		}

		// Get current workflow
		workflow, ok := getWorkflowsCurrent(params)
		if !ok {
			// Current workflow is not set.  Don't play the game
			return "", false, nil
		}

		// get Workflow's map
		wfmap, ok := getWorkflowsWorkflow(params, workflow)
		if !ok {
			return "", false, fmt.Errorf("workflows/map subsection, %s, not in correct format", workflow)
		}

		// Get nextStage from current workflow map
		nextStage, ok := wfmap[stage]
		if !ok {
			return "", false, clearMachineWorkflow(uuid)
		}

		// Get Reboot flag from next Stage
		if sobj, err := Get(nextStage, so); err == nil {
			stage := sobj.(*models.Stage)
			reboot = stage.Reboot
		} else {
			return "", false, fmt.Errorf("Stage, %s, not found", nextStage)
		}
	}

	return nextStage, reboot, nil
}

func nextPxeBootMachine(uuid string) error {
	d, err := session.Machines.GetMachineActions(machines.NewGetMachineActionsParams().WithUUID(strfmt.UUID(uuid)), basicAuth)
	if err != nil {
		return generateError(err, "Failed to fetch actions: %v", uuid)
	}
	actions := d.Payload

	for _, aa := range actions {
		if aa.Command == "nextbootpxe" {
			actionParams := map[string]interface{}{}
			if _, err := session.Machines.PostMachineAction(machines.NewPostMachineActionParams().WithBody(actionParams).WithUUID(strfmt.UUID(uuid)).WithName("nextbootpxe"), basicAuth); err != nil {
				return generateError(err, "Error running action: nextbootpxe")
			}
			break
		}
	}
	return nil
}

func processJobsCommand() *cobra.Command {
	mo := &MachineOps{CommonOps{Name: "machines", SingularName: "machine"}}
	jo := &JobOps{CommonOps{Name: "jobs", SingularName: "job"}}
	so := &StageOps{CommonOps{Name: "stages", SingularName: "stage"}}

	command := &cobra.Command{
		Use:   "processjobs [id]",
		Short: "For the given machine, process pending jobs until done.",
		Long: `
For the provided machine, identified by UUID, process the task list on
that machine until an error occurs or all jobs are complete.  Upon 
completion, optionally wait for additional jobs as specified by
the stage runner wait flag.
`,
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("%v requires at least 1 argument", c.UseLine())
			}
			if len(args) > 1 {
				return fmt.Errorf("%v requires at most 1 arguments", c.UseLine())
			}
			dumpUsage = false

			uuid := args[0]

			var machine *models.Machine
			if obj, err := Get(uuid, mo); err != nil {
				return generateError(err, "Error getting machine")
			} else {
				machine = obj.(*models.Machine)
			}

			fmt.Printf("Processing jobs for %s\n", uuid)

			// Get Current Job and mark it failed if it is running.
			if obj, err := Get(machine.CurrentJob.String(), jo); err == nil {
				job := obj.(*models.Job)
				// If job is running or created, mark it as failed
				if *job.State == "running" || *job.State == "created" {
					markJob(machine.CurrentJob.String(), "failed", jo)
				}
			}

			// Reset Machine information.
			params, err := getMachineParams(uuid)
			if err != nil {
				return generateError(err, "Error getting machine params")
			}

			// Should we use workflows
			version, ok := getWorkflowsVersion(params)
			if ok && version == "1.0" {
				// Get current workflow
				workflow, ok := getWorkflowsCurrent(params)

				// Current workflow isn't set, need to set it
				if !ok {
					workflow, ok = findCurrentWorkflow(params, machine.Stage)
					if ok {
						if err := setMachineWorkflow(uuid, workflow); err != nil {
							return generateError(err, "Failed to set machine's workflow")
						}
					}
				}

				// if workflow is valid, set the stage to the beginning stage for workflow
				if ok {
					newStage, err := getWorkflowsWorkflowFirstStage(params, workflow)
					if err == nil && newStage != machine.Stage {
						if err := setMachineStage(uuid, newStage, mo); err != nil {
							return generateError(err, "Error setting next stage")
						}
					} else if err != nil {
						return generateError(err, "Error finding containing workflow")
					}
				}
			}

			// reset the task index
			setMachineCurrentTask(uuid, -1, mo)

			// Mark machine as runnable
			markMachineRunnable(uuid, mo)

			did_job := false
			for {
				// Wait for machine to be runnable.
				if answer, err := mo.DoWait(machine.UUID.String(), "Runnable", "true", 100000000); err != nil {
					fmt.Printf("Error waiting for machine to be runnable: %v, try again...\n", err)
					time.Sleep(5 * time.Second)
					continue
				} else if answer == "timeout" {
					fmt.Printf("Waiting for machine runnable returned with, %s, trying again.\n", answer)
					continue
				} else if answer == "interrupt" {
					fmt.Printf("User interrupted the wait, exiting ...\n")
					break
				}

				// Create a job for tasks - API Call to get next job on machine
				var job *models.Job
				if obj, err := jo.Create(&models.Job{Machine: machine.UUID}); err != nil {
					fmt.Printf("Error creating a job for machine: %v, continuing\n", err)
					time.Sleep(5 * time.Second)
					continue
				} else {
					// No object, means we don't have work to do.
					if obj == nil {
						if did_job {
							fmt.Println("Jobs finished")
							did_job = false
						}

						// Get the machine's stage. It will tell us what is next.
						// 1. check to see if we should wait for more tasks
						// 2. if no more tasks and no waiting, do a stage change.
						wait := false
						reboot := false
						nextStage := ""
						if obj, err := Get(uuid, mo); err == nil {
							machine = obj.(*models.Machine)

							if sobj, err := Get(machine.Stage, so); err == nil {
								stage := sobj.(*models.Stage)
								wait = stage.RunnerWait

								// We aren't waiting, check to see if we have a new stage
								if !wait {
									if nextStage, reboot, err = getNextStage(uuid, machine.Stage, so); err != nil {
										return generateError(err, "Error getting next stage")
									}
								}
							}
						}

						if wait {
							// Wait for new jobs - XXX: Web socket one day.
							// Create a not equal waiter
							time.Sleep(5 * time.Second)
							continue
						} else {
							if nextStage != "" {
								// Set stage on machine
								if err := setMachineStage(uuid, nextStage, mo); err != nil {
									return generateError(err, "Error setting next stage")
								}

								// Reboot to next bootenv
								if reboot {
									if err := nextPxeBootMachine(uuid); err != nil {
										return generateError(err, "Error setting next pxe boot")
									}

									_, err := exec.Command("reboot").Output()
									if err != nil {
										return generateError(err, "Error rebooting node")
									}
								} else {
									// We didn't reboot, go get tasks
									continue
								}
							}
							break
						}
					}
					job = obj.(*models.Job)
				}
				did_job = true

				// Get the job data
				var list []*models.JobAction
				if resp, err := session.Jobs.GetJobActions(jobs.NewGetJobActionsParams().WithUUID(*job.UUID), basicAuth); err != nil {
					s := fmt.Sprintf("Error loading task content: %v, continuing", err)
					fmt.Printf(s)
					Log(job.UUID, s)
					markJob(job.UUID.String(), "failed", jo)
					continue
				} else {
					list = resp.Payload
				}

				// Mark job as running
				if err := markJob(job.UUID.String(), "running", jo); err != nil {
					markJob(job.UUID.String(), "failed", jo)
					continue
				}
				fmt.Printf("Starting Task: %s (%s)\n", job.Task, job.UUID.String())

				failed := false
				incomplete := false
				reboot := false
				state := "finished"

				for _, action := range list {
					event := &models.Event{Time: strfmt.DateTime(time.Now()), Type: "jobs", Action: "action_start", Key: job.UUID.String(), Object: fmt.Sprintf("Starting task: %s, template: %s", job.Task, *action.Name)}

					if _, err := session.Events.PostEvent(events.NewPostEventParams().WithBody(event), basicAuth); err != nil {
						fmt.Printf("Error posting event: %v\n", err)
					}

					// Excute task
					if *action.Path == "" {
						fmt.Printf("Running Task Template: %s\n", *action.Name)
						failed, incomplete, reboot = runContent(job.UUID, action)
					} else {
						fmt.Printf("Putting Content in place for Task Template: %s\n", *action.Name)
						var s string
						if err := writeStringToFile(*action.Path, *action.Content); err != nil {
							failed = true
							s = fmt.Sprintf("Task Template: %s - Copying contents to %s failed\n%v", *action.Name, *action.Path, err)
						} else {
							s = fmt.Sprintf("Task Template: %s - Copied contents to %s successfully\n", *action.Name, *action.Path)
						}
						fmt.Printf(s)
						Log(job.UUID, s)
					}

					if failed {
						state = "failed"
					} else if incomplete {
						state = "incomplete"
					}

					fmt.Printf("Task Template , %s, %s\n", *action.Name, state)
					event = &models.Event{Time: strfmt.DateTime(time.Now()), Type: "jobs", Action: "action_stop", Key: job.UUID.String(), Object: fmt.Sprintf("Finished task: %s, template: %s, state: %s", job.Task, *action.Name, state)}

					if _, err := session.Events.PostEvent(events.NewPostEventParams().WithBody(event), basicAuth); err != nil {
						fmt.Printf("Error posting event: %v\n", err)
					}

					if failed || incomplete || reboot {
						break
					}
				}

				fmt.Printf("Task: %s %s\n", job.Task, state)
				markJob(job.UUID.String(), state, jo)
				// Loop back and wait for the machine to get marked runnable again

				if reboot {
					_, err := exec.Command("reboot").Output()
					if err != nil {
						Log(job.UUID, "Failed to issue reboot\n")
					}
				}

				// If we failed, should we exit
				if exitOnFailure && failed {
					return fmt.Errorf("Task failed, exiting ...\n")
				}
			}

			return nil
		},
	}
	command.Flags().BoolVar(&exitOnFailure, "exit-on-failure", false, "Exit on failure of a task")

	return command
}
