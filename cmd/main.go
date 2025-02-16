package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "gosplash!!")
	})

	http.ListenAndServe(":8080", mux)
}
