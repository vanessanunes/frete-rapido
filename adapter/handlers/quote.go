package handlers

import (
	"net/http"

	"github.com/vanessanunes/frete-rapido/adapter/postgres/quoterepository"
	"github.com/vanessanunes/frete-rapido/core/domain/integration"

	"github.com/vanessanunes/frete-rapido/core/dto"
)

type Repository struct {
	DB quoterepository.Connection
}

func (repo Repository) Create(w http.ResponseWriter, r *http.Request) {
	quoteRequest, err := dto.FromJSONCreateQuoteRequest(r.Body)
	if err != nil {
		ResponseJson(w, 500, ResponseError{
			Erro: "Ocorreu algum erro ao tentar parsear o Body. Por favor verifique e tente novamente!",
		})
		return
	}

	resp, err := integration.SendRequest(*quoteRequest)
	if err != nil {
		ResponseJson(w, 500, ResponseError{
			Erro: "Erro ao enviar request para serviço de integração.",
		})
		return
	}
	id := repo.DB.Save(resp)
	carriers, err := repo.DB.GetCarriers(id)

	if err != nil {
		ResponseJson(w, 500, ResponseError{
			Erro: "Erro ao buscar informações sobre frete.",
		})
		return
	}

	ResponseJson(w, 200, carriers)
}
