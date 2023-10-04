package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// 2. HandleFunc
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	// 1. Construct http server & register defaultServerMux
	http.ListenAndServe(":9090", nil)
}
