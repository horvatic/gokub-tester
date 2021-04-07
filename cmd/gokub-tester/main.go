package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	started := time.Now()

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(started)
		if duration.Seconds() > 10 {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
		}
	})

	http.HandleFunc("/gokub/health", health)
	http.HandleFunc("/", testReq)
	http.ListenAndServe(":8080", nil)
}

func testReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "echo:, %s!", r.URL.Path[1:])
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
