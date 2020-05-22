package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var ready = true

func status(w http.ResponseWriter, req *http.Request) {
	if ready {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()

	fmt.Fprintf(w, "%s\n", hostname)
	w.WriteHeader(200)
}

func main() {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			<-ticker.C
			ready = !ready
		}
	}()

	http.HandleFunc("/", hello)
	http.HandleFunc("/status", status)
	fmt.Println("Listening on port 8090")
	http.ListenAndServe(":8090", nil)
}
