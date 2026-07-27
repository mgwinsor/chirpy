package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	const fileRootPath = http.Dir(".")

	mux := http.NewServeMux()

	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(fileRootPath)))

	mux.HandleFunc("/healthz", HandlerReadiness)

	svr := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", fileRootPath, port)
	log.Fatal(svr.ListenAndServe())
}
