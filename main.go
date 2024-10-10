package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.String("p", "5442", "specify the listening port")
	folderPath := flag.String("d", ".", "specify the directory to serve")

	flag.Parse()

	fmt.Println("Live @ http://localhost:" + *port)

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*folderPath)))
}
