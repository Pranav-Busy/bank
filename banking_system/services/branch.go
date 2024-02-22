package services

import (
	"github.com/go-pg/pg/v10"
	"github.com/pranav/my/repo/models"
)


func Get_All_Branches_By_Bankid(db *pg.DB,id uint) ([]models.Branch,error){

    ans:=[] models.Branch{}
    err:=db.Model(&ans).Where("bank_id=?",id).Select()
    if err!=nil{
        return nil,err
    }else 
    {
        return ans,nil
    }
}
func Delete_Branch(tx *pg.Tx, db *pg.DB,id models.Branch) error {
    
    _,err:=tx.Model(&id).Where("id=?id").Delete()
    if err!=nil{
        return err
    }else 
    {
        return nil
    }
        
    }



func Get_Branch(db *pg.DB)([]models.Branch,error){

	ans:=[]models.Branch{}

	if err := db.Model(&ans).Select(); err != nil {
      return nil,err
	}
	return ans,nil
}

func Insert_Branch(ans []models.Branch, db*pg.DB) error {
	
	_,err :=db.Model(&ans).Insert()

	if err!=nil{
		return err
	}else 
	{
		return nil
	}
}

func Update_Branch(tx *pg.Tx, db *pg.DB,id models.Branch) error {

    _,err:=tx.Model(&id).Set("name=?name").Where("id=?id").Update()
    if err!=nil{
        return err
    }else 
    {
        return nil
    }
        
    }