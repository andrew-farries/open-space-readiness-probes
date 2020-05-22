package main

import (
	"fmt"
	"net/http"
	"os"
)

func status(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
}

func hello(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()

	fmt.Fprintf(w, "%s\n", hostname)
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/status", status)
	fmt.Println("Listening on port 8090")
	http.ListenAndServe(":8090", nil)
}
