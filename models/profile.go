package models

// Profile represents a set of key/values to use in
// template expansion.
//
// There is one special profile named 'global' that acts
// as a global set of parameters for the system.
//
// These can be assigned to a machine's profile list.
// swagger:model
type Profile struct {
	Validation
	Access
	MetaData
	// The name of the profile.  This must be unique across all
	// profiles.
	//
	// required: true
	Name string
	// A description of this profile.  This can contain any reference
	// information for humans you want associated with the profile.
	Description string
	// Any additional parameters that may be needed to expand templates
	// for BootEnv, as documented by that boot environment's
	// RequiredParams and OptionalParams.
	Params map[string]interface{}
}

func (p *Profile) Prefix() string {
	return "profiles"
}

func (p *Profile) Key() string {
	return p.Name
}

func (p *Profile) Fill() {
	p.Validation.fill()
	p.MetaData.fill()
	if p.Params == nil {
		p.Params = map[string]interface{}{}
	}
}

func (p *Profile) AuthKey() string {
	return p.Key()
}

func (b *Profile) SliceOf() interface{} {
	s := []*Profile{}
	return &s
}

func (b *Profile) ToModels(obj interface{}) []Model {
	items := obj.(*[]*Profile)
	res := make([]Model, len(*items))
	for i, item := range *items {
		res[i] = Model(item)
	}
	return res
}
