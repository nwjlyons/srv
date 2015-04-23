package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

var port = flag.String("p", "8000", "Port to run server on.")

func main() {
	flag.Parse()

	dir, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Serving directory %s on port %s.\n", dir, *port)
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%s", *port),
			http.FileServer(
				http.Dir(dir))))
}
