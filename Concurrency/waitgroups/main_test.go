package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestPrintArgument(t *testing.T) {

	stdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	testData := []string{"apple"}
	var wg sync.WaitGroup

	wg.Add(1)
	go printArgument(testData[0], &wg)

	wg.Wait()
	w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	fmt.Println("output", output)
	os.Stdout = stdout
	if !strings.Contains(output, "apple") {
		t.Errorf("Error as expected apple in the output")
	}

}
