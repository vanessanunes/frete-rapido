package dto

import (
	"encoding/json"
	"io"
)

type CreateQuoteRequest struct {
	Recipient struct {
		Address struct {
			ZipCode string
		}
	}
	Volumes []struct {
		Amount        int
		Category      int
		UnitaryWeight float32 `json:"unitary_weight"`
		Price         float64
		SKU           string
		Height        float32
		Width         float32
		Length        float32
	}
}

func FromJSONCreateQuoteRequest(body io.Reader) (*CreateQuoteRequest, error) {
	createQuoteRequest := CreateQuoteRequest{}
	if err := json.NewDecoder(body).Decode(&createQuoteRequest); err != nil {
		return nil, err
	}
	return &createQuoteRequest, nil
}
