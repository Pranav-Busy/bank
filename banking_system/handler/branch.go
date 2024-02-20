package handler

import (
	"github.com/go-pg/pg/v10"
	"github.com/pranav/my/repo/models"
)


func Getallbranchesbybankid(db *pg.DB,id uint) ([]models.Branch,error){

    ans:=[] models.Branch{}
    err:=db.Model(&ans).Where("bank_id=?",id).Select()
    if err!=nil{
        return nil,err
    }else 
    {
        return ans,nil
    }
}
func DeleteBranch(db *pg.DB,id models.Branch) error {

    _,err:=db.Model(&id).Where("id=?id").Delete()
    if err!=nil{
        return err
    }else 
    {
        return nil
    }
        
    }



func GetBranch(db *pg.DB)([]models.Branch,error){

	ans:=[]models.Branch{}

	if err := db.Model(&ans).Select(); err != nil {
      return nil,err
	}
	return ans,nil
}

func InsertBranch(ans []models.Branch, db*pg.DB) error {
	
	_,err :=db.Model(&ans).Insert()

	if err!=nil{
		return err
	}else 
	{
		return nil
	}
}

func UpdateBranch(db *pg.DB,id models.Branch) error {

    _,err:=db.Model(&id).Set("name=?name").Where("id=?id").Update()
    if err!=nil{
        return err
    }else 
    {
        return nil
    }
        
    }