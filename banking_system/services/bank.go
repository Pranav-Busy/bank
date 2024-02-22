package services

import (
	"github.com/go-pg/pg/v10"
	"github.com/pranav/my/repo/models"
)



func Delete(tx *pg.Tx,db *pg.DB,id models.Bank) error {

	_,err:=tx.Model(&id).Where("id=?id").Delete()
	if err!=nil{
	
		return err
	}else 
	{

		return nil
	}
		
	}



func Details(db *pg.DB)([]models.Bank,error){

	ans:=[]models.Bank{}

	if err := db.Model(&ans).Select(); err != nil {
      return nil,err
	}
	return ans,nil
}

func Insert(ans []models.Bank, db*pg.DB) error {
	
	_,err :=db.Model(&ans).Insert()

	if err!=nil{
		return err
	}else 
	{
		return nil
	}
}


func Update(tx *pg.Tx,db *pg.DB,id models.Bank) error {

	_,err:=tx.Model(&id).Set("name=?name").Where("id=?id").Update()
	if err!=nil{
		return err
	}else 
	{
		return nil
	}
		
	}