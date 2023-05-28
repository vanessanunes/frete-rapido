package handlers

import (
	"log"
	"net/http"
	"strconv"
)

func (repo Repository) Metrics(w http.ResponseWriter, r *http.Request) {
	lastQuote_test := r.URL.Query().Get("last_quote")
	lastQuoteI, err := strconv.Atoi(lastQuote_test)
	if err != nil {
		log.Printf("Erro ao parsear query params de metrics: %v", err)
	}
	metrics, err := repo.DB.GetMetrics(lastQuoteI)
	if err != nil {
		log.Printf("Erro ao gerar metricas: %v", err)
		ResponseJson(w, 500, ResponseError{
			Erro: "Erro ao gerar metricas. Por favor, tente mais tarde.",
		})
		return
	}
	ResponseJson(w, 200, metrics)
}
