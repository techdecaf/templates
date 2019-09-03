package templates

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
)

// Functions struct
type Functions struct {
	Map template.FuncMap
}

func (funcs *Functions) init() error {
	pwd, _ := os.Getwd()

	localFunctions := template.FuncMap{
		"OS":   func() string { return runtime.GOOS },
		"ARCH": func() string { return runtime.GOARCH },
		"CatLines": func(s string) string {
			s = strings.Replace(s, "\r\n", " ", -1)
			return strings.Replace(s, "\n", " ", -1)
		},
		"SplitLines": func(s string) []string {
			return strings.Split(strings.Replace(s, "\r\n", "\n", -1), "\n")
		},
		// "EXEC":      exec,
		"FromSlash": filepath.FromSlash,
		"ToSlash":   filepath.ToSlash,
		"ToTitle":   strings.Title,
		"ToUpper":   strings.ToUpper,
		"ToLower":   strings.ToLower,
		"Replace":   strings.Replace,
		"ReadFile": func(relativePath string) string {
			dat, _ := ioutil.ReadFile(path.Join(pwd, relativePath))
			return string(dat)
		},
		"MkdirAll": func(relativePath string) (err error) {
			err = os.MkdirAll(path.Join(pwd, relativePath), 0700)
			return err
		},
		"Touch": func(relativePath string) (err error) {
			_, err = os.Create(path.Join(pwd, relativePath))
			return err
		},
	}

	funcs.Map = sprig.TxtFuncMap()
	for k, v := range localFunctions {
		funcs.Map[k] = v
	}

	return nil
}
