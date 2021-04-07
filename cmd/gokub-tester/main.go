package main

import (
	"fmt"
	"net/http"
)

func main() {
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
