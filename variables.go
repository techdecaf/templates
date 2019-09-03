package templates

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

// Variable struct
type Variable struct {
	Key   string
	Value string
	// Index int
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
		if err := vars.Set(variable, false); err != nil {
			log.Fatal("variables", fmt.Sprintf("failed to set '%s': %v", variable.Key, err))
		}
	}

	return nil
}

// Set both environment and variable values for use with template expansion
func (vars *Variables) Set(variable Variable, overwrite bool) error {
	key := variable.Key

	val, err := vars.Resolve(variable, overwrite)
	if err != nil {
		return err
	}

	if err := vars.Expander.SetVariable(key, val); err != nil {
		return err
	}
	return nil
}

// Resolve variable values
func (vars *Variables) Resolve(variable Variable, overwrite bool) (val string, err error) {
	reEx := regexp.MustCompile(`^exec\((.*)\)$`)
	// environment
	if env := os.Getenv(variable.Key); env != "" && !overwrite {
		return env, err
	}

	// script values
	expanded, err := vars.Expander.Expand(variable.Value)
	if err != nil {
		return "", err
	}

	if cmd := reEx.FindStringSubmatch(expanded); len(cmd) != 0 {
		return Run(CommandOptions{
			Cmd:       cmd[1],
			UseStdOut: false,
		})
	}

	// default value
	return vars.Expander.Expand(variable.Value)

}
