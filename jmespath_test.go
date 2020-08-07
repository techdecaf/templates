package templates

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestJQSearch(t *testing.T){
  test := Goblin(t)
  test.Describe("given: a package needs to parse a json input", func(){
  test.Describe("when: the JSON is invalid", func(){
    invalidJSON := `{invalid: true}`
    test.It("then: it throws an error", func(){
      var result interface{}
      var err error

      result, err = JMESPath(invalidJSON, "invalid")
      test.Assert(err.Error()).Equal(`invalid character 'i' looking for beginning of object key string`)
      test.Assert(result).Equal("")

    })
  })
	test.Describe("when: the json is valid", func(){
    validJSON := `{
      "goblin": {
        "color":"green", "weight":15, "educated": false, "breakfast": ["muffin", "coffee"]
      }
    }`

	  test.It("then: a specific key can be extracted", func(){
      var result interface{}
      var err error
      // assert
      result, err = JMESPath(validJSON, "goblin.educated")
      test.Assert(result).Equal(false)

      result, err = JMESPath(validJSON, "goblin.breakfast[0]")
      test.Assert(result).Equal("muffin")

      result, err = JMESPath(validJSON, "goblin.color")
      test.Assert(result).Equal("green")

      result, err = JMESPath(validJSON, "goblin.weight")
      test.Assert(result).Equal(float64(15))

      // check for any error result
      test.Assert(err).Equal(nil)
	  })
  })
  })

}