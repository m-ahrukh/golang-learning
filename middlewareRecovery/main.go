package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	// log.Fatal(http.ListenAndServe(":3000", mux))

	log.Fatal(http.ListenAndServe(":3000", recoverPanicMiddleware(mux, true)))
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

func recoverPanicMiddleware(app http.Handler, dev bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()
				log.Println(string(stack))
				if !dev {
					http.Error(w, "Something went wrong ", http.StatusInternalServerError)
					return
				}
				fmt.Fprintf(w, "<h1>pnic: %s</h1><pre>%s</pre>", err, string(stack))
			}
		}()

		app.ServeHTTP(w, r)
	}
}
