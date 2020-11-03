
package main

import (
	"io"

	"regexp"
	"testing"

	pkg "android/soong/java"
)

var t = []testing.InternalTest{

	{"TestAndroidAppImport", pkg.TestAndroidAppImport},

	{"TestAndroidAppImport_DefaultDevCert", pkg.TestAndroidAppImport_DefaultDevCert},

	{"TestAndroidAppImport_Filename", pkg.TestAndroidAppImport_Filename},

	{"TestAndroidAppImport_NoDexPreopt", pkg.TestAndroidAppImport_NoDexPreopt},

	{"TestAndroidAppImport_Presigned", pkg.TestAndroidAppImport_Presigned},

	{"TestAndroidResources", pkg.TestAndroidResources},

	{"TestApp", pkg.TestApp},

	{"TestAppSdkVersion", pkg.TestAppSdkVersion},

	{"TestAppSplits", pkg.TestAppSplits},

	{"TestArchSpecific", pkg.TestArchSpecific},

	{"TestBinary", pkg.TestBinary},

	{"TestCertificates", pkg.TestCertificates},

	{"TestClasspath", pkg.TestClasspath},

	{"TestCollectJavaLibraryPropertiesAddAidlIncludeDirs", pkg.TestCollectJavaLibraryPropertiesAddAidlIncludeDirs},

	{"TestCollectJavaLibraryPropertiesAddJarjarRules", pkg.TestCollectJavaLibraryPropertiesAddJarjarRules},

	{"TestCollectJavaLibraryPropertiesAddLibsDeps", pkg.TestCollectJavaLibraryPropertiesAddLibsDeps},

	{"TestCollectJavaLibraryPropertiesAddScrs", pkg.TestCollectJavaLibraryPropertiesAddScrs},

	{"TestCollectJavaLibraryPropertiesAddStaticLibsDeps", pkg.TestCollectJavaLibraryPropertiesAddStaticLibsDeps},

	{"TestCompilerFlags", pkg.TestCompilerFlags},

	{"TestDefaults", pkg.TestDefaults},

	{"TestDeviceForHost", pkg.TestDeviceForHost},

	{"TestDexpreoptBootJars", pkg.TestDexpreoptBootJars},

	{"TestDexpreoptEnabled", pkg.TestDexpreoptEnabled},

	{"TestDroiddoc", pkg.TestDroiddoc},

	{"TestEmbedNotice", pkg.TestEmbedNotice},

	{"TestExcludeFileGroupInSrcs", pkg.TestExcludeFileGroupInSrcs},

	{"TestGeneratedSources", pkg.TestGeneratedSources},

	{"TestHostForDevice", pkg.TestHostForDevice},

	{"TestInstrumentationTargetOverridden", pkg.TestInstrumentationTargetOverridden},

	{"TestJNIABI", pkg.TestJNIABI},

	{"TestJNIPackaging", pkg.TestJNIPackaging},

	{"TestJarGenrules", pkg.TestJarGenrules},

	{"TestJavaSdkLibrary", pkg.TestJavaSdkLibrary},

	{"TestKapt", pkg.TestKapt},

	{"TestKaptEncodeFlags", pkg.TestKaptEncodeFlags},

	{"TestKotlin", pkg.TestKotlin},

	{"TestNoPlugin", pkg.TestNoPlugin},

	{"TestOverrideAndroidApp", pkg.TestOverrideAndroidApp},

	{"TestPackageNameOverride", pkg.TestPackageNameOverride},

	{"TestPatchModule", pkg.TestPatchModule},

	{"TestPlugin", pkg.TestPlugin},

	{"TestPluginGeneratesApi", pkg.TestPluginGeneratesApi},

	{"TestPrebuilts", pkg.TestPrebuilts},

	{"TestResourceDirs", pkg.TestResourceDirs},

	{"TestResources", pkg.TestResources},

	{"TestSharding", pkg.TestSharding},

	{"TestSimple", pkg.TestSimple},

	{"TestTurbine", pkg.TestTurbine},

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
	return "android/soong/java"
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
