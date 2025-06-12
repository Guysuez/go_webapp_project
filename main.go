package main

import (
	"fmt"
	"net/http"
)

func main() {

	println("Starting webserver on http://localhost:8080/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Got %s request for %s\n", r.Method, r.URL.Path)
		fmt.Fprintln(w, "Hello, DevOps!")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
