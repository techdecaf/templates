package templates

import (
	"bytes"
	"context"
	"io"
	"os"
	"path"
	"strings"

	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/shell"
	"mvdan.cc/sh/syntax"
)

// exists returns whether the given file or directory exists
func exists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}

// CommandOptions - new command runner
type CommandOptions struct {
	Name      string
	Cmd       string
	Dir       string
	Env       []string
	UseStdOut bool
}

// Run - run command
func Run(cmd CommandOptions) (val string, err error) {
	var stdOut io.Writer

	buff := new(bytes.Buffer)
	stdErr := os.Stderr

	if cmd.UseStdOut {
		stdOut = os.Stdout
	} else {
		stdOut = buff
	}

	// set working directory
	pwd, _ := os.Getwd()
	cmd.Dir = path.Join(pwd, cmd.Dir)

	if err := exists(cmd.Dir); err != nil {
		return "", err
	}

	// expand command
	src, err := shell.Expand(cmd.Cmd, os.Getenv)
	if err != nil {
		return "", err
	}

	file, _ := syntax.NewParser().Parse(strings.NewReader(src), cmd.Name)

	runner, _ := interp.New(
		interp.Dir(cmd.Dir),
		interp.StdIO(nil, stdOut, stdErr),
		// interp.WithExecModules(checkInstall),
	)

	// runner.Stdin
	if err := runner.Run(context.TODO(), file); err != nil {
		return "", err
	}

	// if stdErr.String() != "" {
	// 	err = errors.New(stdErr.String())
	// }
	if cmd.UseStdOut {
		return "", err
	}

	return buff.String(), err
}

// func checkInstall(next interp.ExecModule) interp.ExecModule {
// 	return func(ctx context.Context, path string, args []string) error {
// 		if path == "" {
// 			fmt.Printf("Command_Error: %s is not installed\n", args[0])
// 			return interp.ExitStatus(1)
// 		}
// 		return next(ctx, path, args)
// 	}
// }

// import (
// 	"context"
// 	"errors"
// 	"io"
// 	"os"
// 	"path/filepath"
// 	"strings"

// 	"mvdan.cc/sh/expand"
// 	"mvdan.cc/sh/interp"
// 	"mvdan.cc/sh/shell"
// 	"mvdan.cc/sh/syntax"
// )

// // RunCommandOptions is the options for the RunCommand func
// type RunCommandOptions struct {
// 	Command string
// 	Dir     string
// 	Env     []string
// 	Stdin   io.Reader
// 	Stdout  io.Writer
// 	Stderr  io.Writer
// }

// var (
// 	// ErrNilOptions is returned when a nil options is given
// 	ErrNilOptions = errors.New("execext: nil options given")
// )

// // RunCommand runs a shell command
// func RunCommand(ctx context.Context, opts *RunCommandOptions) error {
// 	if opts == nil {
// 		return ErrNilOptions
// 	}

// 	p, err := syntax.NewParser().Parse(strings.NewReader(opts.Command), "")
// 	if err != nil {
// 		return err
// 	}

// 	environ := opts.Env
// 	if len(environ) == 0 {
// 		environ = os.Environ()
// 	}

// 	r, err := interp.New(
// 		interp.Dir(opts.Dir),
// 		interp.Env(expand.ListEnviron(environ...)),

// 		// interp.Module(interp.DefaultExec),
// 		// interp.Module(interp.OpenDevImpls(interp.DefaultOpen)),

// 		interp.StdIO(opts.Stdin, opts.Stdout, opts.Stderr),
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	return r.Run(ctx, p)
// }

// // IsExitError returns true the given error is an exit status error
// func IsExitError(err error) bool {
// 	switch err.(type) {
// 	case interp.ExitStatus, interp.ShellExitStatus:
// 		return true
// 	default:
// 		return false
// 	}
// }

// // Expand is a helper to mvdan.cc/shell.Fields that returns the first field
// // if available.
// func Expand(s string) (string, error) {
// 	s = filepath.ToSlash(s)
// 	s = strings.Replace(s, " ", `\ `, -1)
// 	fields, err := shell.Fields(s, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	if len(fields) > 0 {
// 		return fields[0], nil
// 	}
// 	return "", nil
// }
