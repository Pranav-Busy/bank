package db

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/pranav/my/repo/models"
)

func Create(db *pg.DB) error {
	tx, txerr := db.Begin()
	if txerr != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx.Rollback() 
		return txerr 
	}

	opts := &orm.CreateTableOptions{
		IfNotExists:    true,
		FKConstraints:  true,
	}
	 err := tx.Model(&models.Bank{}).CreateTable(opts)
	if err != nil {
		log.Printf("error while creating table banks, Reason: %v\n", err)
		tx.Rollback() 
		return txerr 
	}
	err1 := tx.Model(&models.Branch{}).CreateTable(opts)
	if err1 != nil {
		log.Printf("error while creating table branches, Reason: %v\n", err1)
		tx.Rollback() 
		return txerr 
	}
	err2 := tx.Model(&models.Customer{}).CreateTable(opts)
	if err1 != nil {
		log.Printf("error while creating table customers, Reason: %v\n", err2)
		tx.Rollback() 
		return txerr 
	}
	err3:= tx.Model(&models.Account{}).CreateTable(opts)
	if err1 != nil {
		log.Printf("error while creating table customers, Reason: %v\n", err3)
		tx.Rollback() 
		return txerr 
	}
	err4:= tx.Model(&models.Transaction{}).CreateTable(opts)
	if err1 != nil {
		log.Printf("error while creating table customers, Reason: %v\n", err4)
		tx.Rollback() 
		return txerr 
	}
	err5:= tx.Model(&models.CustomerAccount{}).CreateTable(opts)
	if err1 != nil {
		log.Printf("error while creating table customers, Reason: %v\n", err5)
		tx.Rollback() 
		return txerr 
	}
	
	
	err6 := tx.Commit() 
	if err != nil {
		log.Printf("error while committing transaction, Reason: %v\n", err6)
		return err3 
	}
	
	return nil
}
