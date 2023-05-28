package domain

type Metrics struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	PriceAvg float64 `json:"price_average"`
	MinPrice float64 `json:"minimum_price"`
	MaxPrice float64 `json:"maximum_price"`
}
