package cli

import (
	"testing"
)

var subnetAddrErrorString string = "Error: GET: subnets: Invalid Address: fred\n\n"
var subnetExpireTimeErrorString string = "Error: GET: subnets: Invalid subnet CIDR: false\n\n"
var subnetShowMissingArgErrorString string = "Error: GET: subnets/ignore: Not Found\n\n"
var subnetExistsMissingIgnoreString string = "Error: GET: subnets/ignore: Not Found\n\n"
var subnetCreateBadJSONErrorString = "Error: Unable to create new subnet: Invalid type passed to subnet create\n\n"
var subnetCreateDuplicateErrorString = "Error: CREATE: subnets/john: already exists\n\n"
var subnetUpdateJohnMissingErrorString string = "Error: GET: subnets/john2: Not Found\n\n"
var subnetPatchJohnMissingErrorString string = "Error: PATCH: subnets/john2: Not Found\n\n"
var subnetDestroyMissingJohnString string = "Error: DELETE: subnets/john: Not Found\n\n"

var subnetInvalidEnabledBooleanListString = "Error: GET: subnets: Enabled must be true or false\n\n"
var subnetRangeIPFailureString string = "Error: PATCH: subnets/john: invalid IP address: cq.98.42.1234\n\n"
var subnetRangeIPBadIpString string = "Error: PATCH: subnets/john: invalid IP address: 192.168.100.500\n\n"
var subnetSubnetCIDRFailureString = "Error: 1111.11.2223.544/66666 is not a valid subnet CIDR\n\n"
var subnetStrategyMacFailureErrorString string = "Error: t5:44:llll:b is not a valid MAC address\n\n"
var subnetLeasetimesIntFailureString string = "Error: 4x5 could not be read as a number\n\n"
var subnetSetIntFailureErrorString string = "Error: 6tl could not be read as a number\n\n"
var subnetGetToNull string = "Error: option 6 does not exist\n\n"

var subnetDefaultListString string = "[]\n"
var subnetEmptyListString string = "[]\n"

var subnetShowNoArgErrorString string = "Error: drpcli subnets show [id] [flags] requires 1 argument\n"
var subnetShowTooManyArgErrorString string = "Error: drpcli subnets show [id] [flags] requires 1 argument\n"

var subnetShowJohnString string = `{
  "ActiveEnd": "192.168.100.100",
  "ActiveLeaseTime": 60,
  "ActiveStart": "192.168.100.20",
  "Available": true,
  "Enabled": false,
  "Errors": [],
  "Name": "john",
  "NextServer": "3.3.3.3",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.255.255.0"
    },
    {
      "Code": 28,
      "Value": "192.168.100.255"
    }
  ],
  "Pickers": [
    "hint",
    "nextFree",
    "mostExpired"
  ],
  "ReadOnly": false,
  "ReservedLeaseTime": 7200,
  "Strategy": "MAC",
  "Subnet": "192.168.100.0/24",
  "Validated": true
}
`

var subnetExistsNoArgErrorString string = "Error: drpcli subnets exists [id] [flags] requires 1 argument"
var subnetExistsTooManyArgErrorString string = "Error: drpcli subnets exists [id] [flags] requires 1 argument"
var subnetExistsIgnoreString string = ""

var subnetCreateNoArgErrorString string = "Error: drpcli subnets create [json] [flags] requires 1 argument\n"
var subnetCreateTooManyArgErrorString string = "Error: drpcli subnets create [json] [flags] requires 1 argument\n"
var subnetCreateBadJSONString = "asdgasdg"

var subnetCreateInputString string = `{
  "Name": "john",
  "ActiveEnd": "192.168.100.100",
  "ActiveStart": "192.168.100.20",
  "ActiveLeaseTime": 60,
  "NextServer": "3.3.3.3",
  "OnlyReservations": false,
  "ReservedLeaseTime": 7200,
  "Subnet": "192.168.100.0/24",
  "Strategy": "MAC"
}
`
var subnetCreateJohnString string = `{
  "ActiveEnd": "192.168.100.100",
  "ActiveLeaseTime": 60,
  "ActiveStart": "192.168.100.20",
  "Available": true,
  "Enabled": false,
  "Errors": [],
  "Name": "john",
  "NextServer": "3.3.3.3",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.255.255.0"
    },
    {
      "Code": 28,
      "Value": "192.168.100.255"
    }
  ],
  "Pickers": [
    "hint",
    "nextFree",
    "mostExpired"
  ],
  "ReadOnly": false,
  "ReservedLeaseTime": 7200,
  "Strategy": "MAC",
  "Subnet": "192.168.100.0/24",
  "Validated": true
}
`

var subnetListBothEnvsString = `[
  {
    "ActiveEnd": "192.168.100.100",
    "ActiveLeaseTime": 60,
    "ActiveStart": "192.168.100.20",
    "Available": true,
    "Enabled": false,
    "Errors": [],
    "Name": "john",
    "NextServer": "3.3.3.3",
    "OnlyReservations": false,
    "Options": [
      {
        "Code": 1,
        "Value": "255.255.255.0"
      },
      {
        "Code": 28,
        "Value": "192.168.100.255"
      }
    ],
    "Pickers": [
      "hint",
      "nextFree",
      "mostExpired"
    ],
    "ReadOnly": false,
    "ReservedLeaseTime": 7200,
    "Strategy": "MAC",
    "Subnet": "192.168.100.0/24",
    "Validated": true
  }
]
`

var subnetUpdateNoArgErrorString string = "Error: drpcli subnets update [id] [json] [flags] requires 2 arguments"
var subnetUpdateTooManyArgErrorString string = "Error: drpcli subnets update [id] [json] [flags] requires 2 arguments"
var subnetUpdateBadJSONString = "asdgasdg"
var subnetUpdateBadJSONErrorString = "Error: Unable to merge objects: json: cannot unmarshal string into Go value of type map[string]interface {}\n\n\n"
var subnetUpdateInputString string = `{
  "Strategy": "NewStrat"
}
`
var subnetUpdateJohnString string = `{
  "ActiveEnd": "192.168.100.100",
  "ActiveLeaseTime": 60,
  "ActiveStart": "192.168.100.20",
  "Available": true,
  "Enabled": false,
  "Errors": [],
  "Name": "john",
  "NextServer": "3.3.3.3",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.255.255.0"
    },
    {
      "Code": 28,
      "Value": "192.168.100.255"
    }
  ],
  "Pickers": [
    "hint",
    "nextFree",
    "mostExpired"
  ],
  "ReadOnly": false,
  "ReservedLeaseTime": 7200,
  "Strategy": "NewStrat",
  "Subnet": "192.168.100.0/24",
  "Validated": true
}
`

var subnetPatchNoArgErrorString string = "Error: drpcli subnets patch [objectJson] [changesJson] [flags] requires 2 arguments"
var subnetPatchTooManyArgErrorString string = "Error: drpcli subnets patch [objectJson] [changesJson] [flags] requires 2 arguments"
var subnetPatchBadPatchJSONString = "asdgasdg"
var subnetPatchBadPatchJSONErrorString = "Error: Unable to parse drpcli subnets patch [objectJson] [changesJson] [flags] JSON asdgasdg\nError: error unmarshaling JSON: json: cannot unmarshal string into Go value of type genmodels.Subnet\n\n"
var subnetPatchBadBaseJSONString = "asdgasdg"
var subnetPatchBadBaseJSONErrorString = "Error: Unable to parse drpcli subnets patch [objectJson] [changesJson] [flags] JSON asdgasdg\nError: error unmarshaling JSON: json: cannot unmarshal string into Go value of type genmodels.Subnet\n\n"
var subnetPatchBaseString string = `{
  "ActiveEnd": "192.168.100.100",
  "ActiveLeaseTime": 60,
  "ActiveStart": "192.168.100.20",
  "Available": true,
  "Enabled": false,
  "Errors": [],
  "Name": "john",
  "NextServer": "3.3.3.3",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.255.255.0"
    },
    {
      "Code": 28,
      "Value": "192.168.100.255"
    }
  ],
  "Pickers": [
    "hint",
    "nextFree",
    "mostExpired"
  ],
  "ReadOnly": false,
  "ReservedLeaseTime": 7200,
  "Strategy": "NewStrat",
  "Subnet": "192.168.100.0/24",
  "Validated": true
}
`
var subnetPatchInputString string = `{
  "Strategy": "bootx64.efi"
}
`
var subnetPatchJohnString string = `{
  "ActiveEnd": "192.168.100.100",
  "ActiveLeaseTime": 60,
  "ActiveStart": "192.168.100.20",
  "Available": true,
  "Enabled": false,
  "Errors": [],
  "Name": "john",
  "NextServer": "3.3.3.3",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.255.255.0"
    },
    {
      "Code": 28,
      "Value": "192.168.100.255"
    }
  ],
  "Pickers": [
    "hint",
    "nextFree",
    "mostExpired"
  ],
  "ReadOnly": false,
  "ReservedLeaseTime": 7200,
  "Strategy": "bootx64.efi",
  "Subnet": "192.168.100.0/24",
  "Validated": true
}
`
var subnetPatchMissingBaseString string = `{
  "ActiveEnd": "192.168.100.100",
  "ActiveLeaseTime": 60,
  "ActiveStart": "192.168.100.20",
  "Enabled": false,
  "Name": "john2",
  "NextServer": "3.3.3.3",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.255.255.0"
    },
    {
      "Code": 28,
      "Value": "192.168.100.255"
    }
  ],
  "Pickers": [
    "hint",
    "nextFree",
    "mostExpired"
  ],
  "ReservedLeaseTime": 7200,
  "Strategy": "bootx64.efi",
  "Subnet": "192.168.100.0/24"
}
`

var subnetDestroyNoArgErrorString string = "Error: drpcli subnets destroy [id] [flags] requires 1 argument"
var subnetDestroyTooManyArgErrorString string = "Error: drpcli subnets destroy [id] [flags] requires 1 argument"
var subnetDestroyJohnString string = "Deleted subnet john\n"

var subnetRangeNoArgErrorString string = "Error: drpcli subnets range [subnetName] [startIP] [endIP] [flags] requires 3 arguments\n"
var subnetRangeTooManyArgErrorString string = "Error: drpcli subnets range [subnetName] [startIP] [endIP] [flags] requires 3 arguments\n"
var subnetRangeIPSuccessString string = `{
  "ActiveEnd": "192.168.100.200",
  "ActiveLeaseTime": 60,
  "ActiveStart": "192.168.100.10",
  "Available": true,
  "Enabled": false,
  "Errors": [],
  "Name": "john",
  "NextServer": "3.3.3.3",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.255.255.0"
    },
    {
      "Code": 28,
      "Value": "192.168.100.255"
    }
  ],
  "Pickers": [
    "hint",
    "nextFree",
    "mostExpired"
  ],
  "ReadOnly": false,
  "ReservedLeaseTime": 7200,
  "Strategy": "NewStrat",
  "Subnet": "192.168.100.0/24",
  "Validated": true
}
`

var subnetSubnetNoArgErrorString string = "Error: drpcli subnets subnet [subnetName] [subnet CIDR] [flags] requires 2 arguments\n"
var subnetSubnetTooManyArgErrorString string = "Error: drpcli subnets subnet [subnetName] [subnet CIDR] [flags] requires 2 arguments\n"
var subnetSubnetCIDRSuccessString = `{
  "ActiveEnd": "192.168.100.200",
  "ActiveLeaseTime": 60,
  "ActiveStart": "192.168.100.10",
  "Available": true,
  "Enabled": false,
  "Errors": [],
  "Name": "john",
  "NextServer": "3.3.3.3",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.192.0.0"
    },
    {
      "Code": 28,
      "Value": "192.191.255.255"
    }
  ],
  "Pickers": [
    "hint",
    "nextFree",
    "mostExpired"
  ],
  "ReadOnly": false,
  "ReservedLeaseTime": 7200,
  "Strategy": "NewStrat",
  "Subnet": "192.168.100.0/10",
  "Validated": true
}
`

var subnetStrategyNoArgErrorString string = "Error: drpcli subnets strategy [subnetName] [MAC] [flags] requires 2 arguments\n"
var subnetStrategyTooManyArgErrorString string = "Error: drpcli subnets strategy [subnetName] [MAC] [flags] requires 2 arguments\n"
var subnetStrategyMacSuccessString string = `{
  "ActiveEnd": "192.168.100.200",
  "ActiveLeaseTime": 60,
  "ActiveStart": "192.168.100.10",
  "Available": true,
  "Enabled": false,
  "Errors": [],
  "Name": "john",
  "NextServer": "3.3.3.3",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.192.0.0"
    },
    {
      "Code": 28,
      "Value": "192.191.255.255"
    }
  ],
  "Pickers": [
    "hint",
    "nextFree",
    "mostExpired"
  ],
  "ReadOnly": false,
  "ReservedLeaseTime": 7200,
  "Strategy": "a3:b3:51:66:7e:11",
  "Subnet": "192.168.100.0/10",
  "Validated": true
}
`

var subnetPickersNoArgErrorString string = "Error: drpcli subnets pickers [subnetName] [list] [flags] requires 2 arguments\n"
var subnetPickersTooManyArgErrorString string = "Error: drpcli subnets pickers [subnetName] [list] [flags] requires 2 arguments\n"
var subnetPickersSuccessString string = `{
  "ActiveEnd": "192.168.100.200",
  "ActiveLeaseTime": 60,
  "ActiveStart": "192.168.100.10",
  "Available": true,
  "Enabled": false,
  "Errors": [],
  "Name": "john",
  "NextServer": "3.3.3.3",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.192.0.0"
    },
    {
      "Code": 28,
      "Value": "192.191.255.255"
    }
  ],
  "Pickers": [
    "none",
    "nextFree",
    "mostExpired"
  ],
  "ReadOnly": false,
  "ReservedLeaseTime": 7200,
  "Strategy": "a3:b3:51:66:7e:11",
  "Subnet": "192.168.100.0/10",
  "Validated": true
}
`

var subnetNextserverNoArgErrorString string = "Error: drpcli subnets nextserver [subnetName] [IP] [flags] requires 2 arguments\n"
var subnetNextserverTooManyArgErrorString string = "Error: drpcli subnets nextserver [subnetName] [IP] [flags] requires 2 arguments\n"
var subnetNextserverIPSuccess string = `{
  "ActiveEnd": "192.168.100.200",
  "ActiveLeaseTime": 60,
  "ActiveStart": "192.168.100.10",
  "Available": true,
  "Enabled": false,
  "Errors": [],
  "Name": "john",
  "NextServer": "1.24.36.16",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.192.0.0"
    },
    {
      "Code": 28,
      "Value": "192.191.255.255"
    }
  ],
  "Pickers": [
    "none",
    "nextFree",
    "mostExpired"
  ],
  "ReadOnly": false,
  "ReservedLeaseTime": 7200,
  "Strategy": "a3:b3:51:66:7e:11",
  "Subnet": "192.168.100.0/10",
  "Validated": true
}
`
var subnetLeasetimesNoArgErrorString string = "Error: drpcli subnets leasetimes [subnetName] [active] [reserved] [flags] requires 3 arguments\n"
var subnetLeasetimesTooManyArgErrorString string = "Error: drpcli subnets leasetimes [subnetName] [active] [reserved] [flags] requires 3 arguments\n"
var subnetLeasetimesSuccessString string = `{
  "ActiveEnd": "192.168.100.200",
  "ActiveLeaseTime": 65,
  "ActiveStart": "192.168.100.10",
  "Available": true,
  "Enabled": false,
  "Errors": [],
  "Name": "john",
  "NextServer": "1.24.36.16",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.192.0.0"
    },
    {
      "Code": 28,
      "Value": "192.191.255.255"
    }
  ],
  "Pickers": [
    "none",
    "nextFree",
    "mostExpired"
  ],
  "ReadOnly": false,
  "ReservedLeaseTime": 7300,
  "Strategy": "a3:b3:51:66:7e:11",
  "Subnet": "192.168.100.0/10",
  "Validated": true
}
`

var subnetSetNoArgErrorString string = "Error: drpcli subnets set [subnetName] option [number] to [value] [flags] requires 5 arguments\n"
var subnetSetTooManyArgErrorString string = "Error: drpcli subnets set [subnetName] option [number] to [value] [flags] requires 5 arguments\n"

var subnetSetTo66 string = `{
  "ActiveEnd": "192.168.100.200",
  "ActiveLeaseTime": 65,
  "ActiveStart": "192.168.100.10",
  "Available": true,
  "Enabled": false,
  "Errors": [],
  "Name": "john",
  "NextServer": "1.24.36.16",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.192.0.0"
    },
    {
      "Code": 28,
      "Value": "192.191.255.255"
    },
    {
      "Code": 6,
      "Value": "66"
    }
  ],
  "Pickers": [
    "none",
    "nextFree",
    "mostExpired"
  ],
  "ReadOnly": false,
  "ReservedLeaseTime": 7300,
  "Strategy": "a3:b3:51:66:7e:11",
  "Subnet": "192.168.100.0/10",
  "Validated": true
}
`
var subnetSetToNull string = `{
  "ActiveEnd": "192.168.100.200",
  "ActiveLeaseTime": 65,
  "ActiveStart": "192.168.100.10",
  "Available": true,
  "Enabled": false,
  "Errors": [],
  "Name": "john",
  "NextServer": "1.24.36.16",
  "OnlyReservations": false,
  "Options": [
    {
      "Code": 1,
      "Value": "255.192.0.0"
    },
    {
      "Code": 28,
      "Value": "192.191.255.255"
    }
  ],
  "Pickers": [
    "none",
    "nextFree",
    "mostExpired"
  ],
  "ReadOnly": false,
  "ReservedLeaseTime": 7300,
  "Strategy": "a3:b3:51:66:7e:11",
  "Subnet": "192.168.100.0/10",
  "Validated": true
}
`
var subnetGetNoArgErrorString string = "Error: drpcli subnets get [subnetName] option [number] [flags] requires 3 arguments\n"
var subnetGetTooManyArgErrorString string = "Error: drpcli subnets get [subnetName] option [number] [flags] requires 3 arguments\n"
var subnetGetTo66 string = "Option 6: 66\n"

func TestSubnetCli(t *testing.T) {
	tests := []CliTest{
		CliTest{true, false, []string{"subnets"}, noStdinString, "Access CLI commands relating to subnets\n", ""},
		CliTest{false, false, []string{"subnets", "list"}, noStdinString, subnetDefaultListString, noErrorString},

		CliTest{true, true, []string{"subnets", "create"}, noStdinString, noContentString, subnetCreateNoArgErrorString},
		CliTest{true, true, []string{"subnets", "create", "john", "john2"}, noStdinString, noContentString, subnetCreateTooManyArgErrorString},
		CliTest{false, true, []string{"subnets", "create", subnetCreateBadJSONString}, noStdinString, noContentString, subnetCreateBadJSONErrorString},
		CliTest{false, false, []string{"subnets", "create", subnetCreateInputString}, noStdinString, subnetCreateJohnString, noErrorString},
		CliTest{false, true, []string{"subnets", "create", subnetCreateInputString}, noStdinString, noContentString, subnetCreateDuplicateErrorString},
		CliTest{false, false, []string{"subnets", "list"}, noStdinString, subnetListBothEnvsString, noErrorString},
		CliTest{false, false, []string{"subnets", "list", "Name=fred"}, noStdinString, subnetEmptyListString, noErrorString},
		CliTest{false, false, []string{"subnets", "list", "Name=john"}, noStdinString, subnetListBothEnvsString, noErrorString},
		CliTest{false, false, []string{"subnets", "list", "Strategy=MAC"}, noStdinString, subnetListBothEnvsString, noErrorString},
		CliTest{false, false, []string{"subnets", "list", "Strategy=false"}, noStdinString, subnetEmptyListString, noErrorString},
		CliTest{false, false, []string{"subnets", "list", "NextServer=3.3.3.3"}, noStdinString, subnetListBothEnvsString, noErrorString},
		CliTest{false, false, []string{"subnets", "list", "NextServer=1.1.1.1"}, noStdinString, subnetEmptyListString, noErrorString},
		CliTest{false, true, []string{"subnets", "list", "NextServer=fred"}, noStdinString, noContentString, subnetAddrErrorString},
		CliTest{false, false, []string{"subnets", "list", "Enabled=false"}, noStdinString, subnetListBothEnvsString, noErrorString},
		CliTest{false, false, []string{"subnets", "list", "Enabled=true"}, noStdinString, subnetEmptyListString, noErrorString},
		CliTest{false, true, []string{"subnets", "list", "Enabled=george"}, noStdinString, noContentString, subnetInvalidEnabledBooleanListString},
		CliTest{false, false, []string{"subnets", "list", "Subnet=192.168.103.0/24"}, noStdinString, subnetEmptyListString, noErrorString},
		CliTest{false, false, []string{"subnets", "list", "Subnet=192.168.100.0/24"}, noStdinString, subnetListBothEnvsString, noErrorString},
		CliTest{false, true, []string{"subnets", "list", "Subnet=false"}, noStdinString, noContentString, subnetExpireTimeErrorString},
		CliTest{true, true, []string{"subnets", "show"}, noStdinString, noContentString, subnetShowNoArgErrorString},
		CliTest{true, true, []string{"subnets", "show", "john", "john2"}, noStdinString, noContentString, subnetShowTooManyArgErrorString},
		CliTest{false, true, []string{"subnets", "show", "ignore"}, noStdinString, noContentString, subnetShowMissingArgErrorString},
		CliTest{false, false, []string{"subnets", "show", "john"}, noStdinString, subnetShowJohnString, noErrorString},

		CliTest{true, true, []string{"subnets", "exists"}, noStdinString, noContentString, subnetExistsNoArgErrorString},
		CliTest{true, true, []string{"subnets", "exists", "john", "john2"}, noStdinString, noContentString, subnetExistsTooManyArgErrorString},
		CliTest{false, false, []string{"subnets", "exists", "john"}, noStdinString, subnetExistsIgnoreString, noErrorString},
		CliTest{false, true, []string{"subnets", "exists", "ignore"}, noStdinString, noContentString, subnetExistsMissingIgnoreString},
		CliTest{true, true, []string{"subnets", "exists", "john", "john2"}, noStdinString, noContentString, subnetExistsTooManyArgErrorString},

		CliTest{true, true, []string{"subnets", "update"}, noStdinString, noContentString, subnetUpdateNoArgErrorString},
		CliTest{true, true, []string{"subnets", "update", "john", "john2", "john3"}, noStdinString, noContentString, subnetUpdateTooManyArgErrorString},
		CliTest{false, true, []string{"subnets", "update", "john", subnetUpdateBadJSONString}, noStdinString, noContentString, subnetUpdateBadJSONErrorString},
		CliTest{false, false, []string{"subnets", "update", "john", subnetUpdateInputString}, noStdinString, subnetUpdateJohnString, noErrorString},
		CliTest{false, true, []string{"subnets", "update", "john2", subnetUpdateInputString}, noStdinString, noContentString, subnetUpdateJohnMissingErrorString},
		CliTest{false, false, []string{"subnets", "show", "john"}, noStdinString, subnetUpdateJohnString, noErrorString},

		CliTest{true, true, []string{"subnets", "patch"}, noStdinString, noContentString, subnetPatchNoArgErrorString},
		CliTest{true, true, []string{"subnets", "patch", "john", "john2", "john3"}, noStdinString, noContentString, subnetPatchTooManyArgErrorString},
		CliTest{false, true, []string{"subnets", "patch", subnetPatchBaseString, subnetPatchBadPatchJSONString}, noStdinString, noContentString, subnetPatchBadPatchJSONErrorString},
		CliTest{false, true, []string{"subnets", "patch", subnetPatchBadBaseJSONString, subnetPatchInputString}, noStdinString, noContentString, subnetPatchBadBaseJSONErrorString},
		CliTest{false, false, []string{"subnets", "patch", subnetPatchBaseString, subnetPatchInputString}, noStdinString, subnetPatchJohnString, noErrorString},
		CliTest{false, true, []string{"subnets", "patch", subnetPatchMissingBaseString, subnetPatchInputString}, noStdinString, noContentString, subnetPatchJohnMissingErrorString},
		CliTest{false, false, []string{"subnets", "show", "john"}, noStdinString, subnetPatchJohnString, noErrorString},

		CliTest{true, true, []string{"subnets", "destroy"}, noStdinString, noContentString, subnetDestroyNoArgErrorString},
		CliTest{true, true, []string{"subnets", "destroy", "john", "june"}, noStdinString, noContentString, subnetDestroyTooManyArgErrorString},
		CliTest{false, false, []string{"subnets", "destroy", "john"}, noStdinString, subnetDestroyJohnString, noErrorString},
		CliTest{false, true, []string{"subnets", "destroy", "john"}, noStdinString, noContentString, subnetDestroyMissingJohnString},
		CliTest{false, false, []string{"subnets", "list"}, noStdinString, subnetDefaultListString, noErrorString},

		CliTest{false, false, []string{"subnets", "create", "-"}, subnetCreateInputString + "\n", subnetCreateJohnString, noErrorString},
		CliTest{false, false, []string{"subnets", "list"}, noStdinString, subnetListBothEnvsString, noErrorString},
		CliTest{false, false, []string{"subnets", "update", "john", "-"}, subnetUpdateInputString + "\n", subnetUpdateJohnString, noErrorString},
		CliTest{false, false, []string{"subnets", "show", "john"}, noStdinString, subnetUpdateJohnString, noErrorString},

		CliTest{true, true, []string{"subnets", "range"}, noStdinString, noContentString, subnetRangeNoArgErrorString},
		CliTest{true, true, []string{"subnets", "range", "john", "1.24.36.7", "1.24.36.16", "1.24.36.16"}, noStdinString, noContentString, subnetRangeTooManyArgErrorString},
		CliTest{false, true, []string{"subnets", "range", "john", "192.168.100.10", "192.168.100.500"}, noStdinString, noContentString, subnetRangeIPBadIpString},
		CliTest{false, true, []string{"subnets", "range", "john", "cq.98.42.1234", "1.24.36.16"}, noStdinString, noContentString, subnetRangeIPFailureString},
		CliTest{false, false, []string{"subnets", "range", "john", "192.168.100.10", "192.168.100.200"}, noStdinString, subnetRangeIPSuccessString, noErrorString},

		CliTest{true, true, []string{"subnets", "subnet"}, noStdinString, noContentString, subnetSubnetNoArgErrorString},
		CliTest{true, true, []string{"subnets", "subnet", "john", "june", "1.24.36.16"}, noStdinString, noContentString, subnetSubnetTooManyArgErrorString},
		CliTest{false, false, []string{"subnets", "subnet", "john", "192.168.100.0/10"}, noStdinString, subnetSubnetCIDRSuccessString, noErrorString},
		CliTest{false, true, []string{"subnets", "subnet", "john", "1111.11.2223.544/66666"}, noStdinString, noContentString, subnetSubnetCIDRFailureString},

		CliTest{true, true, []string{"subnets", "strategy"}, noStdinString, noContentString, subnetStrategyNoArgErrorString},
		CliTest{true, true, []string{"subnets", "strategy", "john", "june", "a3:b3:51:66:7e:11"}, noStdinString, noContentString, subnetStrategyTooManyArgErrorString},
		CliTest{false, false, []string{"subnets", "strategy", "john", "a3:b3:51:66:7e:11"}, noStdinString, subnetStrategyMacSuccessString, noErrorString},
		CliTest{false, true, []string{"subnets", "strategy", "john", "t5:44:llll:b"}, noStdinString, noContentString, subnetStrategyMacFailureErrorString},

		CliTest{true, true, []string{"subnets", "pickers"}, noStdinString, noContentString, subnetPickersNoArgErrorString},
		CliTest{true, true, []string{"subnets", "pickers", "john", "june", "test1,test2,test3"}, noStdinString, noContentString, subnetPickersTooManyArgErrorString},
		CliTest{false, false, []string{"subnets", "pickers", "john", "none,nextFree,mostExpired"}, noStdinString, subnetPickersSuccessString, noErrorString},

		CliTest{true, true, []string{"subnets", "nextserver"}, noStdinString, noContentString, subnetNextserverNoArgErrorString},
		CliTest{true, true, []string{"subnets", "nextserver", "john", "june", "1.24.36.16"}, noStdinString, noContentString, subnetNextserverTooManyArgErrorString},
		CliTest{false, false, []string{"subnets", "nextserver", "john", "1.24.36.16"}, noStdinString, subnetNextserverIPSuccess, noErrorString},

		CliTest{true, true, []string{"subnets", "leasetimes"}, noStdinString, noContentString, subnetLeasetimesNoArgErrorString},
		CliTest{true, true, []string{"subnets", "leasetimes", "john", "june", "32", "55"}, noStdinString, noContentString, subnetLeasetimesTooManyArgErrorString},
		CliTest{false, false, []string{"subnets", "leasetimes", "john", "65", "7300"}, noStdinString, subnetLeasetimesSuccessString, noErrorString},
		CliTest{false, true, []string{"subnets", "leasetimes", "john", "4x5", "55"}, noStdinString, noContentString, subnetLeasetimesIntFailureString},

		CliTest{true, true, []string{"subnets", "set"}, noStdinString, noContentString, subnetSetNoArgErrorString},
		CliTest{true, true, []string{"subnets", "set", "john", "option", "45", "to", "34", "77"}, noStdinString, noContentString, subnetSetTooManyArgErrorString},
		CliTest{true, true, []string{"subnets", "get"}, noStdinString, noContentString, subnetGetNoArgErrorString},
		CliTest{true, true, []string{"subnets", "get", "john", "option", "45", "77"}, noStdinString, noContentString, subnetGetTooManyArgErrorString},
		CliTest{false, true, []string{"subnets", "set", "john", "option", "6tl", "to", "66"}, noStdinString, noContentString, subnetSetIntFailureErrorString},
		CliTest{false, false, []string{"subnets", "set", "john", "option", "6", "to", "66"}, noStdinString, subnetSetTo66, noErrorString},
		CliTest{false, false, []string{"subnets", "get", "john", "option", "6"}, noStdinString, subnetGetTo66, noErrorString},
		CliTest{false, false, []string{"subnets", "set", "john", "option", "6", "to", "null"}, noStdinString, subnetSetToNull, noErrorString},
		CliTest{false, true, []string{"subnets", "get", "john", "option", "6"}, noStdinString, noContentString, subnetGetToNull},

		//End of Helpers

		CliTest{false, false, []string{"subnets", "destroy", "john"}, noStdinString, subnetDestroyJohnString, noErrorString},
		CliTest{false, false, []string{"subnets", "list"}, noStdinString, subnetDefaultListString, noErrorString},
	}

	for _, test := range tests {
		testCli(t, test)
	}

}
