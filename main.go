package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ngfalabella/api-music-golang/handlers"
)

func main() {
	const port = 3030
	addr := fmt.Sprintf(":%d", port)
	mux := http.NewServeMux()

	mux.HandleFunc("/cancion", handlers.HandlerCanciones)

	err := http.ListenAndServe(addr, mux)

	if err != nil {
		log.Fatalln("Ocurrio un error al inciar servidor ", err)
	}
}
