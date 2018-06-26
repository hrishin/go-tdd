package main

import (
	"fmt"
	"io"
	"net/http"
)

// io.Writter has method Write(b []bytes) n int, err Error
func Greeting(writter io.Writer, name string) {
	fmt.Fprintf(writter, "Hello, %s", name)
}

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	Greeting(w, "John")
}

func main() {
	//Greeting(os.Stdout, "Ema")
	http.ListenAndServe(":5000", http.HandlerFunc(GreetingHandler))
}
