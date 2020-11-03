
package main

import (
	"io"

	"os"

	"regexp"
	"testing"

	pkg "android/soong/apex"
)

var t = []testing.InternalTest{

	{"TestApexInProductPartition", pkg.TestApexInProductPartition},

	{"TestApexKeyFromOtherModule", pkg.TestApexKeyFromOtherModule},

	{"TestApexWithExplicitStubsDependency", pkg.TestApexWithExplicitStubsDependency},

	{"TestApexWithShBinary", pkg.TestApexWithShBinary},

	{"TestApexWithStubs", pkg.TestApexWithStubs},

	{"TestApexWithSystemLibsStubs", pkg.TestApexWithSystemLibsStubs},

	{"TestApexWithTarget", pkg.TestApexWithTarget},

	{"TestBasicApex", pkg.TestBasicApex},

	{"TestBasicZipApex", pkg.TestBasicZipApex},

	{"TestFilesInSubDir", pkg.TestFilesInSubDir},

	{"TestHeaderLibsDependency", pkg.TestHeaderLibsDependency},

	{"TestKeys", pkg.TestKeys},

	{"TestMacro", pkg.TestMacro},

	{"TestNonTestApex", pkg.TestNonTestApex},

	{"TestPrebuilt", pkg.TestPrebuilt},

	{"TestPrebuiltFilenameOverride", pkg.TestPrebuiltFilenameOverride},

	{"TestStaticLinking", pkg.TestStaticLinking},

	{"TestTestApex", pkg.TestTestApex},

	{"TestUseVendor", pkg.TestUseVendor},

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
	return "android/soong/apex"
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
