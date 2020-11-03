
package main

import (
	"io"

	"os"

	"regexp"
	"testing"

	pkg "github.com/google/blueprint/proptools"
)

var t = []testing.InternalTest{

	{"TestAppendMatchingProperties", pkg.TestAppendMatchingProperties},

	{"TestAppendProperties", pkg.TestAppendProperties},

	{"TestCloneEmptyProperties", pkg.TestCloneEmptyProperties},

	{"TestCloneProperties", pkg.TestCloneProperties},

	{"TestExtendMatchingProperties", pkg.TestExtendMatchingProperties},

	{"TestExtendProperties", pkg.TestExtendProperties},

	{"TestExternalShellEscaping", pkg.TestExternalShellEscaping},

	{"TestHasTag", pkg.TestHasTag},

	{"TestNinjaEscaping", pkg.TestNinjaEscaping},

	{"TestPropertyIndexesWithTag", pkg.TestPropertyIndexesWithTag},

	{"TestShellEscaping", pkg.TestShellEscaping},

	{"TestTypeEqualProperties", pkg.TestTypeEqualProperties},

	{"TestZeroProperties", pkg.TestZeroProperties},

}

var e = []testing.InternalExample{

}

var matchPat string
var matchRe *regexp.Regexp

type matchString struct{}

func MatchString(pat, str string) (result bool, err error) {
	if matchRe == nil || matchPat != pat {
		matchPat = pat
		matchRe, err = regexp.Compile(matchPat)
		if err != nil {
			return
		}
	}
	return matchRe.MatchString(str), nil
}

func (matchString) MatchString(pat, str string) (bool, error) {
	return MatchString(pat, str)
}

func (matchString) StartCPUProfile(w io.Writer) error {
	panic("shouldn't get here")
}

func (matchString) StopCPUProfile() {
}

func (matchString) WriteHeapProfile(w io.Writer) error {
    panic("shouldn't get here")
}

func (matchString) WriteProfileTo(string, io.Writer, int) error {
    panic("shouldn't get here")
}

func (matchString) ImportPath() string {
	return "github.com/google/blueprint/proptools"
}

func (matchString) StartTestLog(io.Writer) {
	panic("shouldn't get here")
}

func (matchString) StopTestLog() error {
	panic("shouldn't get here")
}

func main() {

	m := testing.MainStart(matchString{}, t, nil, e)


	os.Exit(m.Run())

}
