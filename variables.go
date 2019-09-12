package templates

import (
	"fmt"
	"log"
	"os"
)

// Variable struct
type Variable struct {
	Key         string
	Value       string
	OverrideEnv bool
}

// Variables struct
type Variables struct {
	List     []Variable
	Expander Expander
}

// Init initialize new variables struct
func (vars *Variables) Init() error {
	// load template function helpers
	if err := vars.Expander.Init(); err != nil {
		return err
	}

	// set initial map
	vars.Expander.Variables = make(map[string]interface{})
	for _, variable := range vars.List {
		if err := vars.Set(variable); err != nil {
			log.Fatal("variables", fmt.Sprintf("failed to set '%s': %v", variable.Key, err))
		}
	}

	vars.Expander.Functions.Add("ExpandString", func(str string) string {
		out, _ := vars.Expander.Expand(str)
		return out
	})

	vars.Expander.Functions.Add("ExpandFile", func(file string) string {
		out, _ := vars.Expander.ExpandFile(file)
		return out
	})

	return nil
}

// Set both environment and variable values for use with template expansion
func (vars *Variables) Set(v Variable) error {
	key := v.Key

	val, err := vars.Resolve(v)
	if err != nil {
		return err
	}

	if err := vars.Expander.SetVariable(key, val); err != nil {
		return err
	}
	return nil
}

// Resolve variable values
func (vars *Variables) Resolve(v Variable) (val string, err error) {
	if env := os.Getenv(v.Key); env != "" && !v.OverrideEnv {
		return env, err
	}

	// default value
	return vars.Expander.Expand(v.Value)

}
