package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/vanessanunes/frete-rapido/configs"
)

func OpenConnection() *sql.DB {
	conf := configs.GetDB()
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)
	conn, err := sql.Open(conf.Driver, sc)
	if err != nil {
		log.Panic("Erro ao criar conexão com banco de dados!")
		os.Exit(1)
	}
	if err = conn.Ping(); err != nil {
		log.Println(err)
		conn.Close()
	}

	return conn
}
