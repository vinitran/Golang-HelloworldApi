package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Opened api in 127.0.0.1:8080/")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
