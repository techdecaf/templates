package templates

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestJQSearch(t *testing.T){
  test := Goblin(t)
  test.Describe("given: a package needs to parse a json input", func(){
	test.Describe("when: the json is valid", func(){
    validJSON := `{
      "goblin": {
        "color":"green", "weight":15, "educated": false, "breakfast": ["muffin", "coffee"]
      }
    }`
	  test.It("then: a specific key can be extracted", func(){
		// assert
    test.Assert(JQ(validJSON, "goblin.educated")).Equal(false)
    test.Assert(JQ(validJSON, "goblin.breakfast[0]")).Equal("muffin")
    test.Assert(JQ(validJSON, "goblin.color")).Equal("green")
    test.Assert(JQ(validJSON, "goblin.weight")).Equal(float64(15))
    // test.Assert()
	  })
  })
  })

}