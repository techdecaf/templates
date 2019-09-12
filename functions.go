package templates

import (
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

// Functions struct
type Functions struct {
	Map template.FuncMap
}

func (funcs *Functions) init() error {
	funcs.Map = sprig.TxtFuncMap()

	funcs.Add("OS", func() string { return runtime.GOOS })
	funcs.Add("ARCH", func() string { return runtime.GOARCH })

	funcs.Add("FromSlash", filepath.FromSlash)
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

	funcs.Add("ReadFile", func(file string) string {
		data, _ := ioutil.ReadFile(internal.PathTo(file))
		return string(data)
	})

	funcs.Add("WriteFile", func(data, file string) error {
		return ioutil.WriteFile(internal.PathTo(file), []byte(data), 0700)
	})

	funcs.Add("MkdirAll", func(file string) (err error) {
		err = os.MkdirAll(internal.PathTo(file), 0700)
		return err
	})

	funcs.Add("Touch", func(file string) (err error) {
		_, err = os.Create(internal.PathTo(file))
		return err
	})

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

	return nil
}

// Add functions to the global template functions list.
func (funcs *Functions) Add(key string, action interface{}) {
	funcs.Map[key] = action
}
