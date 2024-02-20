package db

import (
	//"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

func Connect() *pg.DB {

	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "1234",
	})

	if db == nil {
		log.Print("unable to  connect")
		os.Exit(100)
	}

	// log.Print("connection to database was successful")
	//CreateTableProducts(db)
	return db
}
