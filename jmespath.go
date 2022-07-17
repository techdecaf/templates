package templates

import (
	"encoding/json"

	"github.com/ghodss/yaml"
	"github.com/jmespath/go-jmespath"
)

// SearchJSON a jmespath implementation for go
func SearchJSON(input string, search string) (data interface{}, err error) {
	if err := json.Unmarshal(json.RawMessage(input), &data); err != nil {
		return "", err
	}
	result, err := jmespath.Search(search, data)
	if err != nil {
		return "", err
	}

	return result, nil
}

// SearchYAML converts yaml to JSON and runs a jmespath search
func SearchYAML(input string, search string) (data interface{}, err error) {
	rawJSON, err := yaml.YAMLToJSON([]byte(input))
	if err != nil {
		return nil, err
	}
	return SearchJSON(string(rawJSON), search)
}
