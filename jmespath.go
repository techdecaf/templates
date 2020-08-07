package templates

import (
	"encoding/json"

	"github.com/jmespath/go-jmespath"
)

// JMESPath a jmespath implementation for go
func JMESPath(input string, search string) (data interface{}, err error) {
	if err := json.Unmarshal(json.RawMessage(input), &data); err != nil {
		return "", err
	}
	result, err := jmespath.Search(search, data)
	if (err != nil){
		return "", err
	}

	return result, nil
}