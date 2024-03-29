package test

import (
	"testing"

	. "github.com/franela/goblin"
	"github.com/techdecaf/templates"
)

const validYAML = `
goblin:
  color: green
  weight: 15
  educated: false
  breakfast:
    - muffin
    - coffee
`

const validJSON = `{
  "goblin": {
    "color":"green", "weight":15, "educated": false, "breakfast": ["muffin", "coffee"]
  }
}`

func TestJMESPath(t *testing.T) {
	test := Goblin(t)

	test.Describe("given: a package needs to parse a JSON or YAML input", func() {
		test.Describe("when: the input is valid yaml", func() {
			test.It("then: the yaml is converted to JSON", func() {
				// arrange
				result, err := templates.SearchYAML(validYAML, "goblin.breakfast[1]")
				test.Assert(err).Equal(nil)
				test.Assert(result).Equal("coffee")
			})
		})

		test.Describe("when: the JSON is invalid", func() {
			invalidJSON := `{invalid: true}`
			test.It("then: it throws an error", func() {
				var result interface{}
				var err error

				result, err = templates.SearchJSON(invalidJSON, "invalid")
				test.Assert(err.Error()).Equal(`invalid character 'i' looking for beginning of object key string`)
				test.Assert(result).Equal("")

			})
		})
		test.Describe("when: the json is valid", func() {

			test.It("then: a specific key can be extracted", func() {
				var result interface{}
				var err error
				// assert
				result, err = templates.SearchJSON(validJSON, "goblin.educated")
				test.Assert(result).Equal(false)

				result, err = templates.SearchJSON(validJSON, "goblin.breakfast[0]")
				test.Assert(result).Equal("muffin")

				result, err = templates.SearchJSON(validJSON, "goblin.color")
				test.Assert(result).Equal("green")

				result, err = templates.SearchJSON(validJSON, "goblin.weight")
				test.Assert(result).Equal(float64(15))

				// check for any error result
				test.Assert(err).Equal(nil)
			})
		})
	})
}
