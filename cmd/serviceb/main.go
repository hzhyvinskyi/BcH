package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8084"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REQUEST_RECEIVED")

	_, err := w.Write([]byte("Trace!"))
	if err != nil {
		log.Fatalln("FATAL")
	}
}

func main() {
	http.HandleFunc("/trace", handler)

	log.Println("HTTP Server is running on port "+port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("FATAL")
	}
}
