package test

import (
	"os"
	"strings"
	"testing"

	. "github.com/franela/goblin"
	"github.com/techdecaf/templates"
)

func TestVariables(t *testing.T) {
	test := Goblin(t)

	test.Describe("given: a developer needs to set template variables", func() {
		test.Describe("when: a variable already exists in the environment", func() {
			test.It("then: the value should not be overwritten", func() {
				variables := templates.Variables{}
				variables.Init()
				shouldNotOverride := templates.Variable{
					Key:   "TEST_OVERRIDE_VAR",
					Value: "should_not_overwrite",
				}
				os.Setenv(shouldNotOverride.Key, "set_in_env")
				variables.Set(shouldNotOverride)
				// assert
				test.Assert(os.Getenv("TEST_OVERRIDE_VAR")).Equal("set_in_env")
			})
		})

		test.Describe("when: a an explicit override value is set", func() {
			test.It("then: the value should replace what is currently in the env", func() {
				variables := templates.Variables{}
				variables.Init()
				shouldNotOverride := templates.Variable{
					Key:         "TEST_OVERRIDE_VAR",
					Value:       "should_overwrite",
					OverrideEnv: true,
				}
				os.Setenv(shouldNotOverride.Key, "set_in_env")
				variables.Set(shouldNotOverride)
				// assert
				test.Assert(os.Getenv("TEST_OVERRIDE_VAR")).Equal("should_overwrite")
			})
		})

		test.Describe("when: expansion variables are set", func() {
			test.It("then: it should expand variables using EXEC template function", func() {
				variables := templates.Variables{}
				variables.Init()
				variables.Set(templates.Variable{
					Key:   "TEST_EXEC_VAR",
					Value: "{{EXEC `echo go_test`}}",
				})
				// assert
				test.Assert(os.Getenv("TEST_EXEC_VAR")).Equal("go_test")
			})

			test.It("then: it should expand variables using TRY template function", func() {
				variables := templates.Variables{}
				variables.Init()
				variables.Set(templates.Variable{
					Key:   "TEST_TRY_VAR",
					Value: "{{TRY `fails hello` | default `default_value`}}",
				})
				// assert
				test.Assert(os.Getenv("TEST_TRY_VAR")).Equal("default_value")
			})

			test.It("then: it should find and set yaml sub properties using YQ", func() {
				variables := templates.Variables{}
				variables.Init()

				variables.Set(templates.Variable{
					Key: "VALID_YAML",
					// NOTE: validYAML is a const in jmespath_test.go
					Value: validYAML,
				})

				variables.Set(templates.Variable{
					Key:   "TEST_YQ_STRING_VAR",
					Value: "{{YQ `goblin.color` .VALID_YAML }}",
				})
				// assert
				test.Assert(os.Getenv("TEST_YQ_STRING_VAR")).Equal("green")

			})

			test.It("then: it should find and set json sub properties using JQ", func() {
				variables := templates.Variables{}
				variables.Init()
				// strings
				variables.Set(templates.Variable{
					Key:   "TEST_JQ_STRING_VAR",
					Value: "{{JQ `key.sub` `{\"key\": {\"sub\": \"value\"}}` }}",
				})
				// assert
				test.Assert(os.Getenv("TEST_JQ_STRING_VAR")).Equal("value")

				// booleans
				variables.Set(templates.Variable{
					Key:   "TEST_JQ_BOOL_VAR",
					Value: "{{JQ `isTrue` `{\"isTrue\": true}` }}",
				})
				// assert
				test.Assert(os.Getenv("TEST_JQ_BOOL_VAR")).Equal("true")

				// numbers
				variables.Set(templates.Variable{
					Key:   "TEST_JQ_FLOAT_VAR",
					Value: "{{JQ `isTrue` `{\"isTrue\": 32.1337}` }}",
				})
				// assert
				test.Assert(os.Getenv("TEST_JQ_FLOAT_VAR")).Equal("32.1337")
			})

			test.It("then: it should expand to the current working directory", func() {
				variables := templates.Variables{}
				variables.Init()
				// full current working directory
				variables.Set(templates.Variable{
					Key:   "TEST_CWD_VAR",
					Value: "{{PWD}}",
				})
				test.Assert(strings.HasSuffix(os.Getenv("TEST_CWD_VAR"), "/templates/test")).Equal(true)

				// use base to pull out base path only
				variables.Set(templates.Variable{
					Key:   "TEST_CWD_BASE_PATH",
					Value: "{{PWD | base}}",
				})
				test.Assert(os.Getenv("TEST_CWD_BASE_PATH")).Equal("test")
			})

			test.It("then: it should expand variables in files", func() {
				variables := templates.Variables{}
				variables.Init()
				variables.Set(templates.Variable{
					Key:   "BAR",
					Value: "bar",
				})

				// act
				got, err := templates.ExpandFile("data/testfile.txt", variables.Functions)
				if err != nil {
					t.Errorf("failed to expand file %v", err)
				}

				// assert
				test.Assert(strings.Replace(got, "\n", ",", -1)).Equal("FOO=BAR,FOO=BAR")
			})
		})
	})
}

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
