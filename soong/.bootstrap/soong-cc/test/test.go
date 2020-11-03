
package main

import (
	"io"

	"regexp"
	"testing"

	pkg "android/soong/cc"
)

var t = []testing.InternalTest{

	{"TestArchGenruleCmd", pkg.TestArchGenruleCmd},

	{"TestCompilerFlags", pkg.TestCompilerFlags},

	{"TestDataTests", pkg.TestDataTests},

	{"TestDoubleLoadableDepError", pkg.TestDoubleLoadableDepError},

	{"TestDoubleLoadbleDep", pkg.TestDoubleLoadbleDep},

	{"TestExcludeRuntimeLibs", pkg.TestExcludeRuntimeLibs},

	{"TestFuchsiaDeps", pkg.TestFuchsiaDeps},

	{"TestFuchsiaTargetDecl", pkg.TestFuchsiaTargetDecl},

	{"TestGen", pkg.TestGen},

	{"TestLibraryReuse", pkg.TestLibraryReuse},

	{"TestLinkReordering", pkg.TestLinkReordering},

	{"TestLlndkHeaders", pkg.TestLlndkHeaders},

	{"TestPrebuilt", pkg.TestPrebuilt},

	{"TestProto", pkg.TestProto},

	{"TestRecovery", pkg.TestRecovery},

	{"TestRuntimeLibs", pkg.TestRuntimeLibs},

	{"TestRuntimeLibsNoVndk", pkg.TestRuntimeLibsNoVndk},

	{"TestSplitFileExt", pkg.TestSplitFileExt},

	{"TestSplitListForSize", pkg.TestSplitListForSize},

	{"TestStaticDepsOrderWithStubs", pkg.TestStaticDepsOrderWithStubs},

	{"TestStaticExecutable", pkg.TestStaticExecutable},

	{"TestStaticLibDepExport", pkg.TestStaticLibDepExport},

	{"TestStaticLibDepReordering", pkg.TestStaticLibDepReordering},

	{"TestStaticLibDepReorderingWithShared", pkg.TestStaticLibDepReorderingWithShared},

	{"TestVendorModuleUseVndkExt", pkg.TestVendorModuleUseVndkExt},

	{"TestVendorPublicLibraries", pkg.TestVendorPublicLibraries},

	{"TestVendorSrc", pkg.TestVendorSrc},

	{"TestVersionedStubs", pkg.TestVersionedStubs},

	{"TestVndk", pkg.TestVndk},

	{"TestVndkDepError", pkg.TestVndkDepError},

	{"TestVndkExt", pkg.TestVndkExt},

	{"TestVndkExtError", pkg.TestVndkExtError},

	{"TestVndkExtInconsistentSupportSystemProcessError", pkg.TestVndkExtInconsistentSupportSystemProcessError},

	{"TestVndkExtUseVendorLib", pkg.TestVndkExtUseVendorLib},

	{"TestVndkExtVendorAvailableFalseError", pkg.TestVndkExtVendorAvailableFalseError},

	{"TestVndkExtWithoutBoardVndkVersion", pkg.TestVndkExtWithoutBoardVndkVersion},

	{"TestVndkMustNotBeProductSpecific", pkg.TestVndkMustNotBeProductSpecific},

	{"TestVndkSpExtUseVndkError", pkg.TestVndkSpExtUseVndkError},

	{"TestVndkUseVndkExtError", pkg.TestVndkUseVndkExtError},

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
	return "android/soong/cc"
}

func (matchString) StartTestLog(io.Writer) {
	panic("shouldn't get here")
}

func (matchString) StopTestLog() error {
	panic("shouldn't get here")
}

func main() {

	m := testing.MainStart(matchString{}, t, nil, e)


	pkg.TestMain(m)

}
