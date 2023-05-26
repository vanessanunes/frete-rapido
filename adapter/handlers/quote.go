package handlers

import (
	"net/http"

	"github.com/vanessanunes/frete-rapido/core/domain/integration"
	"github.com/vanessanunes/frete-rapido/core/dto"
)

func Create(w http.ResponseWriter, r *http.Request) {
	quoteRequest, err := dto.FromJSONCreateQuoteRequest(r.Body)

	resp := integration.SendRequest(*quoteRequest)

	// guardar isso no banco de dados

	// quotes := Quotes

	if err != nil {
		w.Write([]byte(err.Error()))
		ResponseJson(w, 500, "Houve algum erro na request")
		return
	}

	ResponseJson(w, 200, resp)
}
