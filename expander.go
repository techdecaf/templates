package templates

import (
	"bytes"
	"path"
	"text/template"

	"github.com/techdecaf/templates/internal"
)

// ExpandOptions to use when expanding a string.
type ExpandOptions struct {
	Functions Functions
}

// Expand Template String
func Expand(str string, fn Functions) (string, error) {
	var output bytes.Buffer
	cmdTemplate, err := template.New("cmd").Funcs(fn.Map).Parse(str)
	if err != nil {
		return "", err
	}
	if err := cmdTemplate.Execute(&output, internal.EnvMap()); err != nil {
		return "", err
	}

	return output.String(), err
}

// ExpandFile and return as a string
func ExpandFile(file string, fn Functions) (string, error) {
	var output bytes.Buffer
	var input = internal.PathTo(file)

	fileTemplate, err := template.New(path.Base(input)).Funcs(fn.Map).ParseFiles(input)
	if err != nil {
		return "", err
	}

	env := internal.EnvMap()
	if err := fileTemplate.Execute(&output, env); err != nil {
		return "", err
	}

	return output.String(), err
}
