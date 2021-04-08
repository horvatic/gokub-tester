package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	namespace := os.Getenv("NAMESPACE")
	service := os.Getenv("SERVICE")
	http.HandleFunc(fmt.Sprintf("/%s/%s/%s", namespace, service, "health"), health)
	http.HandleFunc("/health", health)
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
