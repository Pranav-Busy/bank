package services

import (
	"log"
	"errors"
	"github.com/go-pg/pg/v10"
	"github.com/pranav/my/repo/models"
)


func Get_All_Transactions_By_Accountid(db *pg.DB,id uint) ([]models.Transaction,error){

    ans:=[] models.Transaction{}
    err:=db.Model(&ans).Where("account_id=?",id).Select()
    if err!=nil{
        return nil,err
    }else 
    {
        return ans,nil
    }
}

func Delete_Tran(tx *pg.Tx,db *pg.DB,id models.Transaction) (error) {
    _,err:=tx.Model(&id).Where("id = ?id").Delete()
    if err!=nil{
        return err
    }else 
    {
        return nil
    }
        
    }



func Details_Tran(db *pg.DB)([]models.Transaction,error){

	ans:=[]models.Transaction{}

	if err := db.Model(&ans).Select(); err != nil {
      return nil,err
	}
	return ans,nil
}

func Insert_Tran(ans models.Transaction, db*pg.DB) error {
	
	tx, txerr := db.Begin()
	if txerr != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx.Rollback() 
		return txerr 
	}

	if ans.Mode=="deposit"{

		//var helper models.Account
		_,err:=tx.Model(&models.Account{}).Set("balance = balance +?",ans.Amount).Where("id=?",ans.AccountID).Update()
		if err != nil {
			tx.Rollback() 
			return err 
		}
		
	}else if ans.Mode=="withdrawal"{
             
		var helper models.Account

    if err := tx.Model(&helper).Column("balance").Where("id = ?", ans.AccountID).Select(); err != nil {
        tx.Rollback()
        return err
    }

    if helper.Balance < ans.Amount {
		tx.Rollback()
        return errors.New("not possible")
    }

    _, err := tx.Model(&models.Account{}).Set("balance = balance - ?", ans.Amount).Where("id = ?", ans.AccountID).Update()
    if err != nil {
        tx.Rollback()
        return err
    }


	}else if ans.Mode=="transfer"{
                  
		var helper models.Account

    if err := tx.Model(&helper).Column("balance").Where("id = ?", ans.AccountID).Select(); err != nil {
        tx.Rollback()
        return err
    }

    if helper.Balance < ans.Amount {
		tx.Rollback()
        return errors.New("not possible")
    }

    _, err := tx.Model(&models.Account{}).Set("balance = balance - ?", ans.Amount).Where("id = ?", ans.AccountID).Update()
    if err != nil {
        tx.Rollback()
        return err
    }

	_,err1:=tx.Model(&models.Account{}).Set("balance = balance +?",ans.Amount).Where("acc_number=?",ans.Receiver_accnum).Update()
		if err1 != nil {
			tx.Rollback() 
			return err 
		}
		


	}

   if _,err:= tx.Model(&ans).Insert();err!=nil{

	tx.Rollback();
	return err;
   }

    tx.Commit()
	return nil;


}


func Update_Tran(tx *pg.Tx,db *pg.DB,id models.Transaction) error {
  
	
	_,err:=tx.Model(&id).Set("name=?name").Where("id=?id").Update()
	if err!=nil{
		return err
	}else 
	{
		return nil
	}
		
	}