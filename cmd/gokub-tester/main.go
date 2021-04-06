package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", testReq)
	http.ListenAndServe(":8080", nil)
}

func testReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "echo:, %s!", r.URL.Path[1:])
}
