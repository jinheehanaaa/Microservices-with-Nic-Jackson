package main

import (
	"log"
	"net/http"
)

func main() {
	// 2. HandleFunc
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello World")
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	// 1. Construct http server & register defaultServerMux
	http.ListenAndServe(":9090", nil)
}

// MEMO
// ServerMux is responsible for redirecting path
// HandleFunc takes my function and creates http handler from it and adds it to defalt ServerMux
