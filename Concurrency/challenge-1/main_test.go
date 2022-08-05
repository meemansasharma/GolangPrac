package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {

	wg.Add(1)
	go updateMessage("Test data!!", &wg)
	wg.Wait()

	if msg != "Test data!!" {
		t.Errorf("Error msg nor updated")
	}
}

func Test_printMessage(t *testing.T) {
	//save original output stream
	stdout := os.Stdout

	//create a pipe
	r, w, _ := os.Pipe()

	//set the stdout to our write
	os.Stdout = w
	printMessage()

	w.Close()
	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdout
	t.Logf("Logger  %s", output)
	if !strings.Contains(output, "Test data!!") {
		t.Error("Error")
	}

}

func Test_main(t *testing.T) {

	stdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout
	t.Logf("Logger : %s", output)

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Error :Hello, universe! not found")
	}
	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("Error :Hello, cosmos! not found")
	}
	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Error : not found")
	}

}
