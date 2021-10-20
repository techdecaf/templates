package test

import (
	"strings"
	"testing"

	. "github.com/franela/goblin"
	"github.com/techdecaf/templates"
)

func TestGlobMatch(t *testing.T) {
	test := Goblin(t)

	test.Describe("given: a package needs to perform glob pattern matching", func() {
		test.Describe("when: performing a GlobMatch", func() {

			test.It("then: all files matching the pattern are found", func() {
				files, err := templates.GlobMatch("./data/", "**/*.ext")
				// check for any error result

				test.Assert(strings.Join(files, "|")).Equal("glob_match/match_1.ext|glob_match/match_2.ext")
				test.Assert(err).Equal(nil)
			})
		})
	})

	test.Describe("when: performing a file expansion", func() {

		test.It("then: glob patterns can be template expanded", func() {
			funcs := templates.Functions{}
			funcs.Init()
			content, err := templates.ExpandFile("./data/glob_match.test.txt", funcs)

			test.Assert(content).Equal("glob_match.test.txt")
			test.Assert(err).Equal(nil)
		})
	})

}
