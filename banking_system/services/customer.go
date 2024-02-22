package services

import (
	"github.com/go-pg/pg/v10"
	"github.com/pranav/my/repo/models"
)


func Get_All_Customers_By_Branchid(db *pg.DB,id uint) ([]models.Customer,error){

    ans:=[] models.Customer{}
    err:=db.Model(&ans).Where("branch_id=?",id).Select()
    if err!=nil{
        return nil,err
    }else 
    {
        return ans,nil
    }
}


func Delete_Cust(tx *pg.Tx,db *pg.DB,id models.Customer) (error) {

    _,err:=tx.Model(&id).Where("id = ?id").Delete()
    if err!=nil{
        return err
    }else 
    {
        return nil
    }
        
    }



func Details_Cust(db *pg.DB)([]models.Customer,error){

	ans:=[]models.Customer{}

	if err := db.Model(&ans).Select(); err != nil {
      return nil,err
	}
	return ans,nil
}

func Insert_Cust(tx *pg.Tx, ans*[]models.Customer, db *pg.DB) error {
	
	_,err :=tx.Model(ans).Insert()

	if err!=nil{
		return err
	}else 
	{
		return nil
	}
}


func Update_Cust(tx *pg.Tx,db *pg.DB,id models.Customer) error {

	_,err:=tx.Model(&id).Set("name=?name").Where("id=?id").Update()
	if err!=nil{
		return err
	}else 
	{
		return nil
	}
		
	}