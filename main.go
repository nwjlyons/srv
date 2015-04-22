package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.String("p", "8000", "Port to run server on.")

func main() {
	flag.Parse()

	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%s", *port),
			http.FileServer(
				http.Dir("."))))
}
