package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Greet sends a personalized greeting to writer.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler says Hello, world over HTTP.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "Elodie")
	log.Fatal(http.ListenAndServe(":3000", http.HandlerFunc(MyGreeterHandler)))
}
