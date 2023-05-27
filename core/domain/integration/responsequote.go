package integration

import "time"

type ResponseIntegration struct {
	Dispatchers []Dispatchers `json:"dispatchers"`
}

type Dispatchers struct {
	ID                         string   `json:"id"`
	RequestID                  string   `json:"request_id"`
	RegisteredNumberShipper    string   `json:"registered_number_shipper"`
	RegisteredNumberDispatcher string   `json:"registered_number_dispatcher"`
	ZipcodeOrigin              int      `json:"zipcode_origin"`
	Offers                     []Offers `json:"offers"`
}

type Offers struct {
	Offer                int                  `json:"offer"`
	TableReference       string               `json:"table_reference"`
	SimulationType       int                  `json:"simulation_type"`
	Carrier              Carrier              `json:"carrier"`
	Service              string               `json:"service"`
	ServiceCode          string               `json:"service_code"`
	ServiceDescription   string               `json:"service_description"`
	DeliveryTime         DeliveryTime         `json:"delivery_time"`
	Expiration           time.Time            `json:"expiration"`
	CostPrice            float64              `json:"cost_price"`
	FinalPrice           float64              `json:"final_price"`
	Weights              Weights              `json:"weights"`
	OriginalDeliveryTime OriginalDeliveryTime `json:"original_delivery_time"`
}

type Carrier struct {
	Name             string `json:"name"`
	RegisteredNumber string `json:"registered_number"`
	StateInscription string `json:"state_inscription"`
	Logo             string `json:"logo"`
	Reference        int    `json:"reference"`
	CompanyName      string `json:"company_name"`
}

type DeliveryTime struct {
	Days          int    `json:"days"`
	Hours         int    `json:"hours"`
	Minutes       int    `json:"minutes"`
	EstimatedDate string `json:"estimated_date"`
}

type Weights struct {
	Real  float64 `json:"real"`
	Cubed int64   `json:"cubed"`
	Used  int64   `json:"used"`
}

type OriginalDeliveryTime struct {
	Days          int    `json:"days"`
	Hhours        int    `json:"hours"`
	Minutes       int    `json:"minutes"`
	EstimatedDate string `json:"estimated_date"`
}
