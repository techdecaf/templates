package templates

import (
	"encoding/json"
	"log"

	"github.com/jmespath/go-jmespath"
)

// JQ a jmespath implementation for go
func JQ(input string, search string) interface{} {
	var data interface{}
	var err error

	if err := json.Unmarshal(json.RawMessage(input), &data); err != nil {
		log.Fatal(err)
	}
	result, err := jmespath.Search(search, data)
	if (err != nil){
		log.Fatal(err)
	}

	return result
}