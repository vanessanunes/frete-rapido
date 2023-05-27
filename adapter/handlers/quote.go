package handlers

import (
	"net/http"

	"github.com/vanessanunes/frete-rapido/adapter/postgres/quoterepository"
	"github.com/vanessanunes/frete-rapido/core/domain"
	"github.com/vanessanunes/frete-rapido/core/domain/integration"

	"github.com/vanessanunes/frete-rapido/core/dto"
)

type Repository struct {
	DB quoterepository.Connection
}

func (repo Repository) Create(w http.ResponseWriter, r *http.Request) {
	quoteRequest, err := dto.FromJSONCreateQuoteRequest(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		ResponseJson(w, 500, "Ocorreu algum erro ao tentar parsear o Body. Por favor verifique e tente novamente!")
		return
	}

	resp := integration.SendRequest(*quoteRequest)

	repo.DB.Save(resp)

	var carriers []domain.Carrier

	for i := 0; i < len(resp.Dispatchers); i++ {
		offers := resp.Dispatchers[i].Offers
		for j := 0; j < len(offers); j++ {
			carriers = append(carriers, domain.Carrier{
				Name:     offers[j].Carrier.Name,
				Service:  offers[j].Service,
				Deadline: offers[j].DeliveryTime.Days,
				Price:    offers[j].CostPrice,
			})
		}
	}

	quote := domain.Quote{
		Carrier: carriers,
	}

	if err != nil {
		w.Write([]byte(err.Error()))
		ResponseJson(w, 500, "Houve algum erro na request")
		return
	}

	ResponseJson(w, 200, quote)
}
