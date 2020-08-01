package templates

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
)


func TestVariables(t *testing.T){
  test := Goblin(t)

  test.Describe("given: a developer needs to set template variables", func(){
    test.Describe("when: a variable already exists in the environment", func(){
      test.It("then: the value should not be overwritten", func(){
        variables := Variables{}
        variables.Init()
        shouldNotOverride := Variable{
          Key: "TEST_OVERRIDE_VAR",
          Value: "should_not_overwrite",
        }
        os.Setenv(shouldNotOverride.Key, "set_in_env")
        variables.Set(shouldNotOverride)
        // assert
        test.Assert(os.Getenv("TEST_OVERRIDE_VAR")).Equal("set_in_env")	
      })
    })

    test.Describe("when: a an explicit override value is set", func(){
      test.It("then: the value should replace what is currently in the env", func(){
        variables := Variables{}
        variables.Init()
        shouldNotOverride := Variable{
          Key: "TEST_OVERRIDE_VAR",
          Value: "should_overwrite",
          OverrideEnv: true,
        }
        os.Setenv(shouldNotOverride.Key, "set_in_env")
        variables.Set(shouldNotOverride)
        // assert
        test.Assert(os.Getenv("TEST_OVERRIDE_VAR")).Equal("should_overwrite")	
      })
    })

    test.Describe("when: expansion variables are set", func(){
      test.It("then: it should expand variables using EXEC template function", func(){
        variables := Variables{}
        variables.Init()
        variables.Set(Variable{
          Key: "TEST_EXEC_VAR",
          Value: "{{EXEC `echo go_test`}}",
        })
        // assert
        test.Assert(os.Getenv("TEST_EXEC_VAR")).Equal("go_test")
      })

      test.It("then: it should expand variables using TRY template function", func(){
        variables := Variables{}
        variables.Init()
        variables.Set(Variable{
          Key: "TEST_TRY_VAR",
          Value: "{{TRY `fails hello` | default `default_value`}}",
        })
        // assert
        test.Assert(os.Getenv("TEST_TRY_VAR")).Equal("default_value")
      })
    })
  })
  
};



// // EXPANDING A FILE
// func TestFileExpansion(t *testing.T) {
// 	want := "FOO=BAR,FOO=BAR"
// 	var FooVar = Variable{
// 		Key:   "BAR",
// 		Value: "bar",
// 	}

// 	if err := vars.Set(FooVar); err != nil {
// 		t.Errorf("failed to set testVar %v", err)
// 	}

// 	got, err := ExpandFile("tests/testfile.txt", vars.Functions)
// 	if err != nil {
// 		t.Errorf("failed to expand file %v", err)
// 	}

// 	test(t, want, strings.Replace(got, "\n", ",", -1))
// }

// // HELPER FUNCTIONS
// func TestEnv2Map(t *testing.T) {
// 	want := "FOO=BAR"
// 	FooVar := Variable{
// 		Key:         "BAR",
// 		Value:       "FOO=BAR",
// 		OverrideEnv: true,
// 	}

// 	if err := vars.Set(FooVar); err != nil {
// 		t.Errorf("failed to set testVar %v", err)
// 	}

// 	env := internal.EnvMap()

// 	test(t, want, env[FooVar.Key])
// }
