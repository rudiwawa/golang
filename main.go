package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", os.Getenv("NAME"))
	})
	http.ListenAndServe(":8080", nil)
}
