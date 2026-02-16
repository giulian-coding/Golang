package main

import (
	"fmt"
	"net/http"

	"github.com/jboursiquot/go-proverbs"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Proserve!"))
	})
	// PLain Text Endpoint
	http.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, proverbs.Random().Saying)
	})

	// REST API JSON Endpoint

	http.ListenAndServe(":8000", nil)
}
