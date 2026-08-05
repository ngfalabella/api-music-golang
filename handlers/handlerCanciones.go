package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ngfalabella/api-music-golang/utils"
)

func HandlerCanciones(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set(utils.HeaderContentType, utils.ContentTypeJSON)

		err := json.NewEncoder(w).Encode(utils.Canciones)
		if err != nil {
			log.Println("Error respondiendo canciones", err)
		}
	default:
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
	}
}
