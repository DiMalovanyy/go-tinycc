package tinycc

import (
	"os"
	"testing"
)

func getTestSourcesDir() string {
	currdir, _ := os.Getwd()
	return currdir + "/test/"
}

func getTestIncludeDir() string {
	return getTestSourcesDir() + "include/"
}

// Files
const (
	// test0.c
	// Compiling: YES
	// Args:	  NO
	// Payload:	  NO
	// Return:	  SUCCESS
	// Link:	  NO
	WorkingEmptySuccessFile = "test_empty_success.c"
)


// Strings
const (
	WorkingEmptyNoArgsSuccessString = `
		int main(void) {
			return 0;
		}
	`
	WorkingEmptySuccessString = `
		int main(int argc, char **argv) {
			return 0;
		}`
	WorkingStandardLibSucessString = `
		#include <stdlib.h>	
		#include <unistd.h>
		int main(int argc, char **argv) {
			exit(EXIT_SUCCESS);
		}`
	WorkingExternalLibSuccessString = `
		#include <lib0.h>
		int main(int argc, char **argv) {
			int var = sum(2, 2);
			return 0;
		}`
	WorkingInternalLibSuccessString = `
		#include "lib0.h"
		int main(int argc, char **argv){
			int var = sum(2, 2);
			return 0;
		}`
	WorkingNoStdSuccessString = `
		int _start() {
			return 0;
		}
	`
)

// Libs
const (
	LibrarySourceString = `
		int sum_lib(int a, int b) {
			return a + b;
		}`
)

func TestTccContextCreateDelete(t *testing.T) {
	tccContext, err := NewTccContext()
	if err != nil {
		t.Errorf("Could not create tccState: %v", err)
	}
	defer tccContext.DeleteContext()
}

func TestCompileStringBasic(t *testing.T) {
	tccContext, _ := NewTccContext()
	defer tccContext.DeleteContext()

	if err := tccContext.CompileString([]byte(WorkingEmptySuccessString)); err != nil {
		t.Fatal(err)
	}
}

func TestCompileStringFromFile(t *testing.T) {
	tccContext, _ := NewTccContext()
	defer tccContext.DeleteContext()
	sourcePath := getTestSourcesDir() + WorkingEmptySuccessFile
	file, err := os.Open(sourcePath)
	if err != nil {
		t.Fatalf("Could not open file: %s. Error: %v", sourcePath, err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatalf("Could not read file Stat: %s. Error: %v", sourcePath, err)
	}
	fileBuffer := make([]byte, fileInfo.Size())
	_, err = file.Read(fileBuffer)
	if err != nil {
		t.Fatalf("Could not read file to buffer: %s. Error: %v", sourcePath, err)
	}
	if err = tccContext.CompileString(fileBuffer); err != nil {
		t.Fatalf("Could not compile code from file %v. Error: %v", sourcePath, err)
	}
}

func TestCompileStringStandartLibs(t *testing.T) {
	tccContext, _ := NewTccContext()
	defer tccContext.DeleteContext()

	if err := tccContext.CompileString([]byte(WorkingStandardLibSucessString)); err != nil {
		t.Fatal(err)
	}
}

func TestCompileStringExrternalFileLib(t *testing.T) {
	tccContext, _ := NewTccContext()
	defer tccContext.DeleteContext()

	// Note should not compile because the -I<IncludeDir> was not specified
	if err := tccContext.CompileString([]byte(WorkingExternalLibSuccessString)); err != ErrTccCouldNotCompileString {
		t.Fatalf("Code should not compile without define headers include directory")
	}

	if err := tccContext.AddIncludePath(getTestIncludeDir()); err != nil {
		t.Fatal(err)
	}

	// Now should compile
	if err := tccContext.CompileString([]byte(WorkingExternalLibSuccessString)); err != nil {
		t.Fatal(err)
	}
}



