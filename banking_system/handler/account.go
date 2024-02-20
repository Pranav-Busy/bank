package handler

import (
	"github.com/go-pg/pg/v10"
	"github.com/pranav/my/repo/models"
)

func Getallaccountsbybranchid(db *pg.DB,id uint) ([]models.Account,error){

    ans:=[] models.Account{}
    err:=db.Model(&ans).Where("branch_id=?",id).Select()
    if err!=nil{
        return nil,err
    }else 
    {
        return ans,nil
    }
}

func Deleteacc(db *pg.DB,id models.Account) (error) {

    _,err:=db.Model(&id).Where("id = ?", id.ID).Delete()
    if err!=nil{
        return err
    }else 
    {
        return nil
    }
        
    }



func Detailsacc(db *pg.DB)([]models.Account,error){

	ans:=[]models.Account{}

	if err := db.Model(&ans).Select(); err != nil {
      return nil,err
	}
	return ans,nil
}

func Insertacc(ans *models.Account, db*pg.DB) error {
	
	_,err :=db.Model(ans).Insert()

	if err!=nil{
		return err
	}else 
	{
		return nil
	}
}


func Updateacc(db *pg.DB,id models.Account) error {

	_,err:=db.Model(&id).Set("name=?name").Where("id=?id").Update()
	if err!=nil{
		return err
	}else 
	{
		return nil
	}
		
	}