package templates

import (
	"bytes"
	"os"
	"text/template"
)

// Expander - resolve template strings
type Expander struct {
	Variables map[string]interface{}
	// private
	functions Functions
}

// Init - new instance of expander
func (expander *Expander) Init() error {
	// load template function helpers
	if err := expander.functions.init(); err != nil {
		return err
	}
	return nil
}

// SetVariable - set expander variables
func (expander *Expander) SetVariable(key string, val interface{}) error {
	if err := os.Setenv(key, val.(string)); err != nil {
		return err
	}

	expander.Variables[key] = val
	return nil
}

// Expand - Expand Template String
func (expander *Expander) Expand(str string) (string, error) {
	// expand string using templating engine
	var command bytes.Buffer

	cmdTemplate, err := template.New("cmd").Funcs(expander.functions.Map).Parse(str)
	if err != nil {
		return "", err
	}

	if err := cmdTemplate.Execute(&command, expander.Variables); err != nil {
		return "", err
	}

	return command.String(), err
}
