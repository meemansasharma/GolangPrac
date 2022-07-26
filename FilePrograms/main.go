package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if os.Args[1] == "" {
		fmt.Println("Enterfile name")
		os.Exit(1)
	}
	filename := os.Args[1]
	//content, err := ioutil.ReadFile(filename)
	content, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, content)
}
