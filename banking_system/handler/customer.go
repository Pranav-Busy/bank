package handler

import (
	"github.com/go-pg/pg/v10"
	"github.com/pranav/my/repo/models"
)


func Getallcustomersbybranchid(db *pg.DB,id uint) ([]models.Customer,error){

    ans:=[] models.Customer{}
    err:=db.Model(&ans).Where("branch_id=?",id).Select()
    if err!=nil{
        return nil,err
    }else 
    {
        return ans,nil
    }
}


func Deletecust(db *pg.DB,id models.Customer) (error) {

    _,err:=db.Model(&id).Where("id = ?id").Delete()
    if err!=nil{
        return err
    }else 
    {
        return nil
    }
        
    }



func Detailscust(db *pg.DB)([]models.Customer,error){

	ans:=[]models.Customer{}

	if err := db.Model(&ans).Select(); err != nil {
      return nil,err
	}
	return ans,nil
}

func Insertcust(ans *[]models.Customer, db*pg.DB) error {
	
	_,err :=db.Model(ans).Insert()

	if err!=nil{
		return err
	}else 
	{
		return nil
	}
}


func Updatecust(db *pg.DB,id models.Customer) error {

	_,err:=db.Model(&id).Set("name=?name").Where("id=?id").Update()
	if err!=nil{
		return err
	}else 
	{
		return nil
	}
		
	}