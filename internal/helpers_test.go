package internal

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T){
  test := Goblin(t)
  test.Describe("given: an import of the internal package", func(){
    test.Describe("when: calling EnvMap with valid params", func(){
      test.It("then: it converts the current os.Environ to a map", func(){
        os.Setenv("GOBLIN_COLOR", "red")
        env := EnvMap()

        test.Assert(env["GOBLIN_COLOR"]).Equal("red")
      })
    })
  })
}