package quoterepository

import (
	"database/sql"
	"log"

	"github.com/vanessanunes/frete-rapido/core/domain"
	"github.com/vanessanunes/frete-rapido/core/domain/integration"

	_ "github.com/lib/pq"
)

type Connection struct {
	db *sql.DB
}

func ConnectionRepository(db *sql.DB) *Connection {
	return &Connection{db}
}

func (conn Connection) Save(respIntegration integration.ResponseIntegration) (id string) {
	data := respIntegration.Dispatchers
	var dispatchersList []domain.Dispatchers
	var offersList []domain.Offers

	for dispatcher := 0; dispatcher < len(data); dispatcher++ {
		dispatchersList = append(dispatchersList, domain.Dispatchers{
			ID:                         data[dispatcher].ID,
			RequestID:                  data[dispatcher].RequestID,
			RegisteredNumberShipper:    data[dispatcher].RegisteredNumberShipper,
			RegisteredNumberDispatcher: data[dispatcher].RegisteredNumberDispatcher,
			ZipcodeOrigin:              data[dispatcher].ZipcodeOrigin,
		})
		offers := data[dispatcher].Offers
		for offer := 0; offer < len(offers); offer++ {
			offersList = append(offersList, domain.Offers{
				Offer:                             offers[offer].Offer,
				TableReference:                    offers[offer].TableReference,
				SimulationType:                    offers[offer].SimulationType,
				Service:                           offers[offer].Service,
				DeliveryTimeDays:                  offers[offer].DeliveryTime.Days,
				DeliveryTimeEstimatedDate:         offers[offer].DeliveryTime.EstimatedDate,
				Expiration:                        offers[offer].Expiration,
				CostPrice:                         offers[offer].CostPrice,
				FinalPrice:                        offers[offer].FinalPrice,
				WeightsReal:                       offers[offer].Weights.Real,
				WeightsUsed:                       offers[offer].Weights.Used,
				OriginalDeliveryTimeDays:          offers[offer].OriginalDeliveryTime.Days,
				OriginalDeliveryTimeEstimatedDate: offers[offer].OriginalDeliveryTime.EstimatedDate,
				Carrier: domain.CarrierM{
					Name:             offers[offer].Carrier.Name,
					RegisteredNumber: offers[offer].Carrier.RegisteredNumber,
					StateInscription: offers[offer].Carrier.StateInscription,
					Logo:             offers[offer].Carrier.Logo,
					Reference:        offers[offer].Carrier.Reference,
					CompanyName:      offers[offer].Carrier.CompanyName,
				},
			})
		}
		id = conn.SaveDispatcher(dispatchersList)
		conn.SaveOffer(offersList, id)
		return id
	}
	return id
}

func (conn *Connection) SaveDispatcher(dispatchers []domain.Dispatchers) (id string) {
	query := `INSERT INTO dispatcher (id, request_id, registered_number_shipper, registered_number_dispatcher, zipcode_origin) VALUES ($1, $2, $3, $4, $5) RETURNING ID`
	dispatcher := dispatchers[0]
	row := conn.db.QueryRow(query, &dispatcher.ID, &dispatcher.RequestID, &dispatcher.RegisteredNumberShipper, &dispatcher.RegisteredNumberDispatcher, &dispatcher.ZipcodeOrigin).Scan(&id)
	if row != nil {
		log.Printf("Erro ao salvar informações de cotação: %v", row)
		return
	}
	return id
}

func (conn *Connection) SaveOffer(offers []domain.Offers, id string) {
	sql := `INSERT INTO offer (id_dispatcher, offer, table_reference, simulation_type, service, delivery_time_days, delivery_time_estimated_date, expiration, cost_price, final_price, weights_real, weights_used, original_delivery_time_days, original_delivery_time_estimated_date) 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id;`
	for col := 0; col < len(offers); col++ {
		idOffer := 0
		conn.db.QueryRow(sql, id, offers[col].Offer, offers[col].TableReference, offers[col].SimulationType, offers[col].Service, offers[col].DeliveryTimeDays, offers[col].DeliveryTimeEstimatedDate, offers[col].Expiration, offers[col].CostPrice, offers[col].FinalPrice, offers[col].WeightsReal, offers[col].WeightsUsed, offers[col].OriginalDeliveryTimeDays, offers[col].OriginalDeliveryTimeEstimatedDate).Scan(&idOffer)
		carrier := domain.CarrierM(offers[col].Carrier)
		carrier.IdOffer = idOffer
		conn.SaveCarrier(carrier)
	}
}

func (conn *Connection) SaveCarrier(carrier domain.CarrierM) {
	sql := `INSERT INTO carrier (id_offer, "name", registered_number, state_inscription, logo, reference, company_name) VALUES($1, $2, $3, $4, $5, $6, $7);`
	row := conn.db.QueryRow(sql, &carrier.IdOffer, &carrier.Name, &carrier.RegisteredNumber, &carrier.StateInscription, &carrier.Logo, &carrier.Reference, &carrier.CompanyName)
	if row != nil {
		log.Printf("Erro ao salvar informações de serviço: %v", row)
	}
}
