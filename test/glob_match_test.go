package test

import (
	"strings"
	"testing"

	. "github.com/franela/goblin"
	"github.com/techdecaf/templates"
	"github.com/techdecaf/templates/internal"
)

func TestGlobMatch(t *testing.T) {
	test := Goblin(t)

	test.Describe("given: a package needs to perform glob pattern matching", func() {
		test.Describe("when: performing a GlobMatch", func() {

			test.It("then: all files matching the pattern are found", func() {
				files, err := templates.GlobMatch("./data/", "**/*.ext")
				// check for any error result

				expected := []string{
					internal.PathTo("./data/glob_match/match_1.ext"),
					internal.PathTo("./data/glob_match/match_2.ext"),
				}
				test.Assert(strings.Join(files, "|")).Equal(strings.Join(expected, "|"))
				test.Assert(err).Equal(nil)
			})
		})
	})

	test.Describe("when: performing a file expansion", func() {

		test.It("then: glob patterns can be template expanded", func() {
			funcs := templates.Functions{}
			funcs.Init()
			content, err := templates.ExpandFile("./data/glob_match.test.txt", funcs)

			test.Assert(content).Equal(internal.PathTo("./data/glob_match.test.txt"))
			test.Assert(err).Equal(nil)
		})
	})

}
