package templates

import (
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T){
  test := Goblin(t)
  test.Describe("given: a tempalate", func(){
    test.Describe("when: functions are required", func(){
      test.It("then: ", func(){
        // arrange

        // act

        // assert

      })
    })
  })

}