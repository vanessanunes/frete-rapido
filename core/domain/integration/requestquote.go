package integration

import "github.com/vanessanunes/frete-rapido/configs"

type Response struct {
	Shipper            Shipper              `json:"shipper"`
	Recipient          Recipient            `json:"recipient"`
	DispatchersRequest []DispatchersRequest `json:"dispatchers"`
	Channel            string               `json:"channel"`
	Filter             int32                `json:"filter"`
	Limit              int32                `json:"limit"`
	Identification     string               `json:"identification"`
	Reverse            bool                 `json:"reverse"`
	SimulationType     []int                `json:"simulation_type"`
	Returns            Returns              `json:"returns"`
}

type Shipper struct {
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}

type Recipient struct {
	Type             int    `json:"type"`
	RegisteredNumber string `json:"registered_number"`
	StateInscription string `json:"state_inscription"`
	Country          string `json:"country"`
	Zipcode          int32  `json:"zipcode"`
}

type DispatchersRequest struct {
	RegisteredNumber string    `json:"registered_number"`
	Zipcode          int32     `json:"zipcode"`
	TotalPrice       float64   `json:"total_price"`
	Volumes          []Volumes `json:"volumes"`
}

type Volumes struct {
	Amount        int     `json:"amount"`
	AmountVolumes int     `json:"amount_volumes"`
	Category      string  `json:"category"`
	SKU           string  `json:"sku"`
	Tag           string  `json:"tag"`
	Description   string  `json:"description"`
	Height        float32 `json:"height"`
	Width         float32 `json:"width"`
	Length        float32 `json:"length"`
	UnitaryPrice  float64 `json:"unitary_price"`
	UnitaryWeight float32 `json:"unitary_weight" default1:"0.1"`
	Consolidate   bool    `json:"consolidate"`
	Overlaid      bool    `json:"overlaid"`
	Rotate        bool    `json:"rotate"`
}

type Returns struct {
	Composition  bool `json:"composition"`
	Volumes      bool `json:"volumes"`
	AppliedRules bool `json:"applied_rules"`
}

func NewShipper() Shipper {
	api := configs.GetServer()
	return Shipper{
		RegisteredNumber: api.DispatcherCNPJ,
		Token:            api.KeyToken,
		PlatformCode:     api.KeyPlataformCode,
	}
}
