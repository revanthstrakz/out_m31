
package main

import (
	"io"

	"os"

	"regexp"
	"testing"

	pkg "android/soong/android"
)

var t = []testing.InternalTest{

	{"TestConsistentNamespaceNames", pkg.TestConsistentNamespaceNames},

	{"TestDeclaringNamespaceInNonAndroidBpFile", pkg.TestDeclaringNamespaceInNonAndroidBpFile},

	{"TestDependingOnBlueprintModuleInRootNamespace", pkg.TestDependingOnBlueprintModuleInRootNamespace},

	{"TestDependingOnModuleByFullyQualifiedReference", pkg.TestDependingOnModuleByFullyQualifiedReference},

	{"TestDependingOnModuleInImportedNamespace", pkg.TestDependingOnModuleInImportedNamespace},

	{"TestDependingOnModuleInNonImportedNamespace", pkg.TestDependingOnModuleInNonImportedNamespace},

	{"TestDependingOnModuleInRootNamespace", pkg.TestDependingOnModuleInRootNamespace},

	{"TestDependingOnModuleInSameNamespace", pkg.TestDependingOnModuleInSameNamespace},

	{"TestDirectorySortedPaths", pkg.TestDirectorySortedPaths},

	{"TestExpand", pkg.TestExpand},

	{"TestFilterArchStruct", pkg.TestFilterArchStruct},

	{"TestFilterList", pkg.TestFilterList},

	{"TestFirstUniqueStrings", pkg.TestFirstUniqueStrings},

	{"TestImplicitlyImportRootNamespace", pkg.TestImplicitlyImportRootNamespace},

	{"TestImportingNonexistentNamespace", pkg.TestImportingNonexistentNamespace},

	{"TestInList", pkg.TestInList},

	{"TestIndexList", pkg.TestIndexList},

	{"TestJoinWithPrefix", pkg.TestJoinWithPrefix},

	{"TestLastUniqueStrings", pkg.TestLastUniqueStrings},

	{"TestMaybeRel", pkg.TestMaybeRel},

	{"TestMissingVendorConfig", pkg.TestMissingVendorConfig},

	{"TestModulesDoReceiveParentNamespace", pkg.TestModulesDoReceiveParentNamespace},

	{"TestNamespaceImportsNotTransitive", pkg.TestNamespaceImportsNotTransitive},

	{"TestNamespaceNotAtTopOfFile", pkg.TestNamespaceNotAtTopOfFile},

	{"TestNamespacesDontInheritParentNamespaces", pkg.TestNamespacesDontInheritParentNamespaces},

	{"TestNeverallow", pkg.TestNeverallow},

	{"TestNewCustomOnceKey", pkg.TestNewCustomOnceKey},

	{"TestNewOnceKey", pkg.TestNewOnceKey},

	{"TestOncePerReentrant", pkg.TestOncePerReentrant},

	{"TestOncePer_Get", pkg.TestOncePer_Get},

	{"TestOncePer_Get_panic", pkg.TestOncePer_Get_panic},

	{"TestOncePer_Get_wait", pkg.TestOncePer_Get_wait},

	{"TestOncePer_Once", pkg.TestOncePer_Once},

	{"TestOncePer_Once2StringSlice", pkg.TestOncePer_Once2StringSlice},

	{"TestOncePer_OnceStringSlice", pkg.TestOncePer_OnceStringSlice},

	{"TestOncePer_Once_wait", pkg.TestOncePer_Once_wait},

	{"TestOptionalPath", pkg.TestOptionalPath},

	{"TestPathDepsMutator", pkg.TestPathDepsMutator},

	{"TestPathForModuleInstall", pkg.TestPathForModuleInstall},

	{"TestPathForModuleSrc", pkg.TestPathForModuleSrc},

	{"TestPathForSource", pkg.TestPathForSource},

	{"TestPathsForModuleSrc", pkg.TestPathsForModuleSrc},

	{"TestPathsForModuleSrc_AllowMissingDependencies", pkg.TestPathsForModuleSrc_AllowMissingDependencies},

	{"TestPrebuiltEtcAndroidMk", pkg.TestPrebuiltEtcAndroidMk},

	{"TestPrebuiltEtcGlob", pkg.TestPrebuiltEtcGlob},

	{"TestPrebuiltEtcHost", pkg.TestPrebuiltEtcHost},

	{"TestPrebuiltEtcOutputPath", pkg.TestPrebuiltEtcOutputPath},

	{"TestPrebuiltEtcVariants", pkg.TestPrebuiltEtcVariants},

	{"TestPrebuiltUserShareHostInstallDirPath", pkg.TestPrebuiltUserShareHostInstallDirPath},

	{"TestPrebuiltUserShareInstallDirPath", pkg.TestPrebuiltUserShareInstallDirPath},

	{"TestPrebuilts", pkg.TestPrebuilts},

	{"TestPrefixInList", pkg.TestPrefixInList},

	{"TestPrintfIntoProperty", pkg.TestPrintfIntoProperty},

	{"TestProductConfigAnnotations", pkg.TestProductConfigAnnotations},

	{"TestRemoveFromList", pkg.TestRemoveFromList},

	{"TestRemoveListFromList", pkg.TestRemoveListFromList},

	{"TestRename", pkg.TestRename},

	{"TestRuleBuilder", pkg.TestRuleBuilder},

	{"TestRuleBuilder_Build", pkg.TestRuleBuilder_Build},

	{"TestSameNameInTwoNamespaces", pkg.TestSameNameInTwoNamespaces},

	{"TestSearchOrder", pkg.TestSearchOrder},

	{"TestTwoModulesWithSameNameInSameNamespace", pkg.TestTwoModulesWithSameNameInSameNamespace},

	{"TestTwoNamepacesInSameDir", pkg.TestTwoNamepacesInSameDir},

	{"TestTwoNamespacesCanImportEachOther", pkg.TestTwoNamespacesCanImportEachOther},

	{"TestValidateConfigAnnotations", pkg.TestValidateConfigAnnotations},

	{"TestValidatePath", pkg.TestValidatePath},

	{"TestValidateSafePath", pkg.TestValidateSafePath},

	{"TestVtsConfig", pkg.TestVtsConfig},

	{"Test_Shard", pkg.Test_Shard},

}

var e = []testing.InternalExample{

	
		{"OutputPath_FileInSameDir", pkg.ExampleOutputPath_FileInSameDir, "out/system/framework/boot.art out/system/framework/oat/arm/boot.vdex\nboot.art oat/arm/boot.vdex\n", false},
	

	
		{"OutputPath_ReplaceExtension", pkg.ExampleOutputPath_ReplaceExtension, "out/system/framework/boot.art out/system/framework/boot.oat\nboot.art boot.oat\n", false},
	

	
		{"RuleBuilder", pkg.ExampleRuleBuilder, "commands: \"ld a.o b.o -o out/linked && echo success\"\ntools: [\"ld\"]\ninputs: [\"a.o\" \"b.o\"]\noutputs: [\"out/linked\"]\n", false},
	

	

	
		{"RuleBuilderCommand_Flag", pkg.ExampleRuleBuilderCommand_Flag, "ls -l\n", false},
	

	
		{"RuleBuilderCommand_FlagForEachArg", pkg.ExampleRuleBuilderCommand_FlagForEachArg, "ls --sort=time --sort=size\n", false},
	

	
		{"RuleBuilderCommand_FlagForEachInput", pkg.ExampleRuleBuilderCommand_FlagForEachInput, "turbine --classpath a.jar --classpath b.jar\n", false},
	

	
		{"RuleBuilderCommand_FlagWithArg", pkg.ExampleRuleBuilderCommand_FlagWithArg, "ls --sort=time\n", false},
	

	
		{"RuleBuilderCommand_FlagWithInput", pkg.ExampleRuleBuilderCommand_FlagWithInput, "java -classpath=a\n", false},
	

	
		{"RuleBuilderCommand_FlagWithInputList", pkg.ExampleRuleBuilderCommand_FlagWithInputList, "java -classpath=a.jar:b.jar\n", false},
	

	
		{"RuleBuilderCommand_FlagWithList", pkg.ExampleRuleBuilderCommand_FlagWithList, "ls --sort=time,size\n", false},
	

	
		{"RuleBuilderCommand_Flags", pkg.ExampleRuleBuilderCommand_Flags, "ls -l -a\n", false},
	

	
		{"RuleBuilder_DeleteTemporaryFiles", pkg.ExampleRuleBuilder_DeleteTemporaryFiles, "commands: \"cp a out/b && cp out/b out/c && rm -f out/b\"\ntools: [\"cp\"]\ninputs: [\"a\"]\noutputs: [\"out/c\"]\n", false},
	

	
		{"RuleBuilder_Installs", pkg.ExampleRuleBuilder_Installs, "rule.Installs().String() = \"out/linked:/bin/linked out/linked:/sbin/linked\"\n", false},
	

	
		{"RuleBuilder_Temporary", pkg.ExampleRuleBuilder_Temporary, "commands: \"cp a out/b && cp out/b out/c\"\ntools: [\"cp\"]\ninputs: [\"a\"]\noutputs: [\"out/c\"]\n", false},
	

	
		{"CopyOf", pkg.ExampleCopyOf, "a = [\"-1\" \"2\" \"3\"]\nb = [\"1\" \"2\" \"3\"]\n", false},
	

	
		{"CopyOf_append", pkg.ExampleCopyOf_append, "Without CopyOf:\na = [\"foo\"]\nb = [\"foo\" \"baz\"]\nc = [\"foo\" \"baz\"]\nWith CopyOf:\na = [\"foo\"]\nb = [\"foo\" \"bar\"]\nc = [\"foo\" \"baz\"]\n", false},
	

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
	return "android/soong/android"
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
