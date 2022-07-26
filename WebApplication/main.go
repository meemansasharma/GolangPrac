package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/getGreeting", func(w http.ResponseWriter, r *http.Request) {

		n, err := fmt.Fprintf(w, " Hello World ")
		if err != nil {
			fmt.Print("Error : ", err)
		}
		fmt.Printf("The number of bytes written are %v", n)
	})

	http.ListenAndServe(":8080", nil)
}
