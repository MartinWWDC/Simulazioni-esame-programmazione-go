package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestFiltro(t *testing.T) {
	RunTest(t, []string{}, "66 1.73\n", "22.052190183434128\nnormale\n")
	RunTest(t, []string{}, "74 2.0\n", "18.5\nnormale\n")
	RunTest(t, []string{}, "51 1.71\n", "17.441263978660103\nsotto\n")
	RunTest(t, []string{}, "100 2.0\n", "25\nsopra\n")
}

func RunTest(t *testing.T, args []string, stdin string, stdout string) { //si possono fare slice per stdin e stdout ma attento ai \n finali, se si vuole testare anche il tempo bisognerebbe fare controlli aggiuntivi
	
	os.Args = append([]string{""},args...)
	
	readStdin, writeStdin, _ := os.Pipe() //pipe written to writeStdin can be read from readStdin
	readStdin, os.Stdin = os.Stdin, readStdin //now we can write program input to writeStdin
	
	readStdout, writeStdout, _ := os.Pipe() //pipe written to writeStdout can be read from readStdout
	writeStdout, os.Stdout = os.Stdout, writeStdout //now we can read program output from readStdout

	fmt.Fprint(writeStdin, stdin)

	main()

	//revert correct stdin and stdout files
	readStdin, os.Stdin = os.Stdin, readStdin
	writeStdout, os.Stdout = os.Stdout, writeStdout
	readStdin.Close()
	writeStdin.Close()
	writeStdout.Close()

	result, _ := io.ReadAll(readStdout)

	readStdout.Close()
	
	if stdout != string(result) {
		t.Error(fmt.Sprintf("\nArgs: %s\nInput: %s\nExpected result:\n-------\n%s\n-------\nGiven result:\n-------\n%s\n-------\n", strings.Join(args, " "), stdin, stdout, string(result)))
	}
}
