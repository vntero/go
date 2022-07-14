package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, young world!\n")
	}

	http.HandleFunc("/hello", helloHandler)

	log.Println("Listening for requests at http://localhost8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
