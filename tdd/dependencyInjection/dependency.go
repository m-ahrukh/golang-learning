package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	// Greet(os.Stdout, "Mahrukh Babar")
	log.Fatal(http.ListenAndServe(":3000", http.HandlerFunc(MyGreetHandler)))
}
