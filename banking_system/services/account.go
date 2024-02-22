package services

import (
	"github.com/go-pg/pg/v10"
	"github.com/pranav/my/repo/models"
)

func Get_All_Accounts_By_Branchid(db *pg.DB,id uint) ([]models.Account,error){

    ans:=[] models.Account{}
    err:=db.Model(&ans).Where("branch_id=?",id).Select()
    if err!=nil{
        return nil,err
    }else 
    {
        return ans,nil
    }
}

func Delete_Acc(tx *pg.Tx,db *pg.DB,id models.Account) (error) {

    _,err:=tx.Model(&id).Where("id = ?", id.ID).Delete()
    if err!=nil{
        return err
    }else 
    {
        return nil
    }
        
    }



func Details_Acc(db *pg.DB)([]models.Account,error){

	ans:=[]models.Account{}

	if err := db.Model(&ans).Select(); err != nil {
      return nil,err
	}
	return ans,nil
}

func Insert_Acc(tx *pg.Tx,ans *models.Account, db*pg.DB) error {
	
	_,err :=tx.Model(ans).Insert()

	if err!=nil{
		return err
	}else 
	{
		return nil
	}
}


func Update_Acc(tx *pg.Tx,db *pg.DB,id models.Account) error {

	if _,err:=tx.Model(&id).Set("name=?name").Where("id=?id").Update();err!=nil{
		return err
	}else 
	{
		return nil
	}
		
	}