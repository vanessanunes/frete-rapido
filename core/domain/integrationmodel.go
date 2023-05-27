package domain

import "time"

type Dispatchers struct {
	ID                         string `json:"id"`
	RequestID                  string `json:"request_id"`
	RegisteredNumberShipper    string `json:"registered_number_shipper"`
	RegisteredNumberDispatcher string `json:"registered_number_dispatcher"`
	ZipcodeOrigin              int    `json:"zipcode_origin"`
}

type Offers struct {
	Offer                             int       `json:"offer"`
	TableReference                    string    `json:"table_reference"`
	SimulationType                    int       `json:"simulation_type"`
	Service                           string    `json:"service"`
	DeliveryTimeDays                  int       `json:"days"`
	DeliveryTimeEstimatedDate         string    `json:"estimated_date"`
	Expiration                        time.Time `json:"expiration"`
	CostPrice                         float64   `json:"cost_price"`
	FinalPrice                        float64   `json:"final_price"`
	WeightsReal                       float64   `json:"weights_real"`
	WeightsUsed                       int64     `json:"weights_used"`
	OriginalDeliveryTimeDays          int       `json:"original_dalivery_days"`
	OriginalDeliveryTimeEstimatedDate string    `json:"original_dalivery_estimated_date"`
	Carrier                           CarrierM
}

type CarrierM struct {
	IdOffer          int    `json:"offer"`
	Name             string `json:"name"`
	RegisteredNumber string `json:"registered_number"`
	StateInscription string `json:"state_inscription"`
	Logo             string `json:"logo"`
	Reference        int    `json:"reference"`
	CompanyName      string `json:"company_name"`
}
