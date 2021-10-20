package templates

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/techdecaf/templates/internal"
)

// Functions template function map
type Functions struct {
	Map template.FuncMap
}

// Add functions to the global template functions list.
func (funcs *Functions) Add(key string, action interface{}) {
	funcs.Map[key] = action
}

// Init adds all default functions to the Map
func (funcs *Functions) Init() error {
	funcs.Map = sprig.TxtFuncMap()

	funcs.Add("OS", func() string { return runtime.GOOS })
	funcs.Add("ARCH", func() string { return runtime.GOARCH })
	funcs.Add("PWD", func() string {
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		return path
	})

	// string manipulation

	funcs.Add("ToSlash", filepath.ToSlash)
	funcs.Add("ToTitle", strings.Title)
	funcs.Add("ToUpper", strings.ToUpper)
	funcs.Add("ToLower", strings.ToLower)
	funcs.Add("Replace", strings.Replace)

	funcs.Add("CatLines", func(s string) string {
		s = strings.Replace(s, "\r\n", " ", -1)
		return strings.Replace(s, "\n", " ", -1)
	})

	funcs.Add("SplitLines", func(s string) []string {
		return strings.Split(strings.Replace(s, "\r\n", "\n", -1), "\n")
	})

	// file system operations
	funcs.Add("FromSlash", filepath.FromSlash)

	funcs.Add("ReadFile", func(file string) string {
		data, _ := ioutil.ReadFile(internal.PathTo(file))
		return string(data)
	})

	funcs.Add("WriteFile", internal.WriteFile)

	funcs.Add("MkdirAll", func(file string) (err error) {
		err = os.MkdirAll(internal.PathTo(file), 0700)
		return err
	})

	funcs.Add("RemoveAll", func(files ...string) (err error) {
		for _, file := range files {
			if err = os.RemoveAll(internal.PathTo(file)); err != nil {
				return err
			}
		}
		return err
	})

	funcs.Add("Expand", func(str string) string {
		out, _ := Expand(str, *funcs)
		return out
	})

	funcs.Add("Touch", func(file string) (err error) {
		_, err = os.Create(internal.PathTo(file))
		return err
	})

	// execution functions

	funcs.Add("EXEC", func(cmd string) string {
		output, err := Run(CommandOptions{
			Cmd:        cmd,
			UseStdOut:  false,
			TrimOutput: true,
		})

		if err != nil {
			log.Fatal(err)
		}

		return output
	})

	funcs.Add("TRY", func(cmd string) string {
		output, err := Run(CommandOptions{
			Cmd:        cmd,
			UseStdOut:  false,
			TrimOutput: true,
		})

		if err != nil {
			return ""
		}

		return output
	})

	funcs.Add("ExpandFile", func(file string) string {
		out, _ := ExpandFile(file, *funcs)
		return out
	})

	funcs.Add("GlobMatch", func(path, pattern string) []string {
		if out, err := GlobMatch(path, pattern); err != nil {
			log.Fatalln(fmt.Sprintf("GlobMatch failed, %v", err))
			return nil
		} else {
			return out
		}
	})

	// serialization functions
	funcs.Add("JQ", func(search string, input string) interface{} {
		out, err := SearchJSON(input, search)
		if err != nil {
			log.Fatal(fmt.Sprintf("JQ failed, %v", err))
		}
		return out
	})

	funcs.Add("YQ", func(search string, input string) interface{} {
		out, err := SearchYAML(input, search)
		if err != nil {
			log.Fatal(fmt.Sprintf("YQ failed, %v", err))
		}
		return out
	})

	return nil
}
