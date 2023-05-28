package quoterepository

import (
	"fmt"
	"log"

	"github.com/vanessanunes/frete-rapido/core/domain"
)

func (conn *Connection) GetMetrics(lastQuote int) (metrics []domain.Metrics, err error) {
	query := `select c.name, count(c.name) as quantity, avg(o.final_price) as price_avg, min(o.final_price) as min_price, max(o.final_price) as max_price
	from offer o
	inner join carrier c on c.id_offer = o.id
	group by c.name 
	order by c.name desc `
	if lastQuote != 0 {
		query += fmt.Sprintf("limit %d", lastQuote)
	}
	rows, err := conn.db.Query(query)
	if err != nil {
		log.Println(err)
		return metrics, err
	}
	var metric domain.Metrics
	for rows.Next() {
		err = rows.Scan(&metric.Name, &metric.Quantity, &metric.PriceAvg, &metric.MinPrice, &metric.MaxPrice)
		if err != nil {
			log.Println(err)
		}
		metrics = append(metrics, metric)
	}
	return
}
