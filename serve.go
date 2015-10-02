package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8080", "The port on which to serve requests.")
	flag.Parse()

	log.Printf("Serving on port %s\n", *port)
	panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir("."))))
}
