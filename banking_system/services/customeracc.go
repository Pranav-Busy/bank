package services

import (
	//"log"

	"github.com/go-pg/pg/v10"
	"github.com/pranav/my/repo/models"
)


func Insert_Ca(tx *pg.Tx,id []models.CustomerAccount) error{
	
	_,err:=tx.Model(&id).Insert()

	if err!=nil{
		return err;
	}

	return nil;

}

func Get_Cust_By_Accid(db *pg.DB,id uint)([]*models.Customer,error){

	var ans models.Account
	getErr := db.Model(&ans).
		Relation("Customer").
		Where("id=?",id).
		Select()


	if getErr != nil {
		//log.Fatal(getErr)
		return nil,getErr
	}

	return ans.Customer,nil
}

func Get_Acc_By_Cid(db *pg.DB,id uint)([]*models.Account,error){

	var ans models.Customer
	getErr := db.Model(&ans).
		Relation("Account").
		Where("id=?",id).
		Select()


	if getErr != nil {
		//log.Fatal(getErr)
		return nil,getErr
	}

	return ans.Account,nil
}

