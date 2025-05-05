package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, This is A DOCKER Learner! Able to push an optimized docker image")
}

func main() {
	fmt.Println("started server...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
