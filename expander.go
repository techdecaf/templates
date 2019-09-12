package templates

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"

	"github.com/techdecaf/templates/internal"
)

// Expander - resolve template strings
type Expander struct {
	Variables map[string]interface{}
	// private
	Functions Functions
}

// Init - new instance of expander
func (expander *Expander) Init() error {
	// load template function helpers
	if err := expander.Functions.init(); err != nil {
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
	var output bytes.Buffer

	cmdTemplate, err := template.New("cmd").Funcs(expander.Functions.Map).Parse(str)
	if err != nil {
		return "", err
	}

	if err := cmdTemplate.Execute(&output, expander.Variables); err != nil {
		return "", err
	}

	return output.String(), err
}

// ExpandFile and return as a string
func (expander *Expander) ExpandFile(file string) (string, error) {
	var output bytes.Buffer
	var input = internal.PathTo(file)
	fmt.Println(input)

	fileTemplate, err := template.New(path.Base(input)).Funcs(expander.Functions.Map).ParseFiles(input)
	if err != nil {
		return "", err
	}

	if err := fileTemplate.Execute(&output, expander.Variables); err != nil {
		return "", err
	}

	return output.String(), err
}

// WriteFile and return as a string
func (expander *Expander) WriteFile(data, file string) error {
	return ioutil.WriteFile(internal.PathTo(file), []byte(data), 0700)
}
