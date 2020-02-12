package main

import (
	"fmt"
	"net/http"
	"os"
)

const defaultGreeting = "World"
const greetingFormat = "Hello %s!"
const address = ":8080"

func main() {
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = defaultGreeting
	}
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, greetingFormat, greeting)
	})
	http.ListenAndServe(address, nil)
}
