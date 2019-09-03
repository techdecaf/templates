package templates

import (
	"fmt"
	"os"
	"testing"
)

var (
	vars          = Variables{}
	ShouldExecVar = Variable{
		Key:   "EXEC_VAR",
		Value: "{{EXEC `echo hello_world`}}",
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
		fmt.Printf("[TEST][PASSED][%s][got = %v]\n", t.Name(), got)
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
