package templates

import (
	"os"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/techdecaf/templates/internal"
)

// GlobMatch
func GlobMatch(dir string, pattern string) ([]string, error) {
	_path := internal.PathTo(dir)
	files, err := doublestar.Glob(os.DirFS(_path), pattern)

	out := make([]string, len(files))
	for i, file := range files {
		out[i] = internal.PathTo(_path + "/" + file)
	}
	return out, err
}
