package integration

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/vanessanunes/frete-rapido/configs"
	"github.com/vanessanunes/frete-rapido/core/dto"
)

func SendRequest(quoteRequest dto.CreateQuoteRequest) ResponseIntegration {
	api := configs.GetServer()

	shipper := NewShipper()
	recipient := Recipient{
		Type:    api.RecipientType,
		Country: api.RecipientCountry,
		Zipcode: api.DispatcherZipcode,
	}

	volumesRequest := quoteRequest.Volumes
	var volumeSend []Volumes
	for i := 0; i < len(volumesRequest); i++ {
		convertCategory := strconv.Itoa(volumesRequest[i].Category)
		volumeSend = append(volumeSend, Volumes{
			Amount:        volumesRequest[i].Amount,
			Category:      convertCategory,
			UnitaryWeight: volumesRequest[i].UnitaryWeight,
			UnitaryPrice:  volumesRequest[i].Price,
			SKU:           volumesRequest[i].SKU,
			Height:        volumesRequest[i].Height,
			Width:         volumesRequest[i].Width,
			Length:        volumesRequest[i].Length,
		})
	}

	var dispatchers []DispatchersRequest
	dispatchers = append(dispatchers, DispatchersRequest{
		RegisteredNumber: api.DispatcherCNPJ,
		Zipcode:          api.DispatcherZipcode,
		Volumes:          volumeSend,
	})

	returns := Returns{
		Composition:  false,
		Volumes:      false,
		AppliedRules: false,
	}

	response := Response{
		Shipper:            shipper,
		Recipient:          recipient,
		DispatchersRequest: dispatchers,
		SimulationType:     []int{0},
		Returns:            returns,
	}

	json_data, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(api.BaseUrl, "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var res ResponseIntegration

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}

	return res

}
