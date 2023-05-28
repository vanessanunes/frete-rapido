package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseError struct {
	Erro string `json:"message"`
}

func ResponseJson(w http.ResponseWriter, statusCode int, resp any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println("Houve algum problema para imprimir a resposta.")
	}
}
