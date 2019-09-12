package templates

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/techdecaf/templates/internal"
)

var vars = Variables{}

var (
	ShouldExecVar = Variable{
		Key:   "EXEC_VAR",
		Value: "{{EXEC `echo hello_world`}}",
	}

	ShouldTryVar = Variable{
		Key:   "TRY_VAR",
		Value: "{{TRY `fails hello` | default `default`}}",
	}

	ShouldOverrideVar = Variable{
		Key:         "OVERRIDE_VAR",
		Value:       "overwritten",
		OverrideEnv: true,
	}

	ShouldNotOverrideVar = Variable{
		Key:   "OVERRIDE_VAR",
		Value: "does_not_overwrite",
	}
)

func test(t *testing.T, want, got interface{}) {
	if got != want {
		t.Errorf("%v() = %q, want %q", t.Name(), got, want)
	} else {
		fmt.Printf("[TEST][PASSED][%s]\n", t.Name())
		fmt.Printf("[want = %v]\n", want)
		fmt.Printf("[got  = %v]\n", got)
		fmt.Println()
	}
}

func TestTemplateExpansion(t *testing.T) {
	want := "hello_world"

	if err := vars.Init(); err != nil {
		t.Errorf("failed to init vars %v", err)
	}

	if err := vars.Set(ShouldExecVar); err != nil {
		t.Errorf("failed to set testVar %v", err)
	}

	test(t, want, os.Getenv(ShouldExecVar.Key))
}

func TestTryFunc(t *testing.T) {
	want := "default"

	if err := vars.Init(); err != nil {
		t.Errorf("failed to init vars %v", err)
	}

	if err := vars.Set(ShouldTryVar); err != nil {
		t.Errorf("failed to set testVar %v", err)
	}

	test(t, want, os.Getenv(ShouldTryVar.Key))
}

func TestEnvResolution(t *testing.T) {
	if err := vars.Init(); err != nil {
		t.Errorf("failed to init vars %v", err)
	}

	// set new env variable
	if err := vars.Set(ShouldNotOverrideVar); err != nil {
		t.Errorf("failed to set testVar %v", err)
	}

	// validate set
	test(t, ShouldNotOverrideVar.Value, os.Getenv(ShouldNotOverrideVar.Key))

	// override variable
	if err := vars.Set(ShouldOverrideVar); err != nil {
		t.Errorf("failed to set testVar %v", err)
	}

	// ensure overwritten
	test(t, ShouldOverrideVar.Value, os.Getenv(ShouldNotOverrideVar.Key))

	// set non overriding variable again
	if err := vars.Set(ShouldNotOverrideVar); err != nil {
		t.Errorf("failed to set testVar %v", err)
	}

	// should not be overwritten
	test(t, ShouldOverrideVar.Value, os.Getenv(ShouldNotOverrideVar.Key))
}

// EXPANDING A FILE
func TestFileExpansion(t *testing.T) {
	want := "FOO=BAR,FOO=BAR"
	var FooVar = Variable{
		Key:   "BAR",
		Value: "bar",
	}

	if err := vars.Set(FooVar); err != nil {
		t.Errorf("failed to set testVar %v", err)
	}

	got, err := ExpandFile("tests/testfile.txt", vars.Functions)
	if err != nil {
		t.Errorf("failed to expand file %v", err)
	}

	test(t, want, strings.Replace(got, "\n", ",", -1))
}

// HELPER FUNCTIONS
func TestEnv2Map(t *testing.T) {
	want := "FOO=BAR"
	FooVar := Variable{
		Key:         "BAR",
		Value:       "FOO=BAR",
		OverrideEnv: true,
	}

	if err := vars.Set(FooVar); err != nil {
		t.Errorf("failed to set testVar %v", err)
	}

	env := internal.EnvMap()

	test(t, want, env[FooVar.Key])
}
