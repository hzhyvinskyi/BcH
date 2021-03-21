package main

import (
	"log"
	"net/http"
)

const port = ":8086"

func handler(w http.ResponseWriter, r *http.Request) {
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
