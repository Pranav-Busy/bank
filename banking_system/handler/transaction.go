package handler

import (
	"github.com/go-pg/pg/v10"
	"github.com/pranav/my/repo/models"
)


func Getalltransactionsbyaccountid(db *pg.DB,id uint) ([]models.Transaction,error){

    ans:=[] models.Transaction{}
    err:=db.Model(&ans).Where("account_id=?",id).Select()
    if err!=nil{
        return nil,err
    }else 
    {
        return ans,nil
    }
}

func Deletetran(db *pg.DB,id models.Transaction) (error) {

    _,err:=db.Model(&id).Where("id = ?id").Delete()
    if err!=nil{
        return err
    }else 
    {
        return nil
    }
        
    }



func Detailstran(db *pg.DB)([]models.Transaction,error){

	ans:=[]models.Transaction{}

	if err := db.Model(&ans).Select(); err != nil {
      return nil,err
	}
	return ans,nil
}

func Inserttran(ans []models.Transaction, db*pg.DB) error {
	
	_,err :=db.Model(&ans).Insert()

	if err!=nil{
		return err
	}else 
	{
		return nil
	}
}


func Updatetran(db *pg.DB,id models.Transaction) error {

	_,err:=db.Model(&id).Set("name=?name").Where("id=?id").Update()
	if err!=nil{
		return err
	}else 
	{
		return nil
	}
		
	}