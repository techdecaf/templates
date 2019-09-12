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
	List      []Variable
	Functions Functions
}

// Init initialize new variables struct
func (vars *Variables) Init() error {
	// load template function helpers
	if err := vars.Functions.init(); err != nil {
		return err
	}

	for _, variable := range vars.List {
		if err := vars.Set(variable); err != nil {
			log.Fatal("variables", fmt.Sprintf("failed to set '%s': %v\n", variable.Key, err))
		}
	}

	return nil
}

// Set both environment and variable values for use with template expansion
func (vars *Variables) Set(v Variable) error {

	val, err := vars.Resolve(v)
	if err != nil {
		return err
	}

	if err := os.Setenv(v.Key, val); err != nil {
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
	return Expand(v.Value, vars.Functions)

}
