package tinycc

import (
	"os"
	"testing"
)

const (
	// test0.c
	// Compiling: YES
	// Args:	  NO
	// Payload:	  NO
	// Return:	  SUCCESS
	// Link:	  NO
	WorkingEmptySuccess = "test_empty_success.c"

	// test1.c
	// Compiling: YES
	// ARGS:	  NO
	// Payload:   
	//		print: "Hello go-tinycc\n"
	// Return:	  SUCESS
	WorkingPayloadSuccess = "test1.c"
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

	source := `
		int main(int argc, char **argv) {
			return 0;
		}`
	if err := tccContext.CompileString([]byte(source)); err != nil {
		t.Fatal(err)
	}
}

func TestCompileStringFromFile(t *testing.T) {
	tccContext, _ := NewTccContext()
	defer tccContext.DeleteContext()
	currdir, _ := os.Getwd()
	sourcePath := currdir + "/test/" + WorkingEmptySuccess
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

func TestCompileStringExternalLibs(t *testing.T) {
	tccContext, _ := NewTccContext()
	defer tccContext.DeleteContext()

	source := `
		#include <stdlib.h>	
		#include <unistd.h>
		int main(int argc, char **argv) {
			exit(EXIT_SUCCESS);
		}`
	if err := tccContext.CompileString([]byte(source)); err != nil {
		t.Fatal(err)
	}
}


