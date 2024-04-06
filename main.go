package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(724)
		w.Write([]byte("hello world!"))
	})

	if err := http.ListenAndServe(":7240", nil); err != nil {
		fmt.Printf("error starting http server: %v", err)
	}
}
