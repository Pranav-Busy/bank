package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/pranav/my/repo/models"
	"github.com/pranav/my/repo/db"
	"github.com/pranav/my/repo/services"

)
var pg_db *pg.DB=db.Connect()
func InsertDetails(c *gin.Context) {
	var ans []models.Bank
	c.ShouldBindJSON(&ans)
	err := services.Insert(ans, pg_db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry data could not be entered in bank table"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data was inserted successfully"})
	}
}

func BankDetails(c *gin.Context) {
	ans, err := services.Details(pg_db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of bank couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}
}
func DeleteBank(c *gin.Context) {
	var id models.Bank
	c.ShouldBindJSON(&id)
	tx, txerr := pg_db.Begin()
	if txerr != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx.Rollback()
		return
	}
	err := services.Delete(tx,pg_db, id)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry bank could not be deleted", "error": err})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "bank deleted successfully"})
	}
}

func UpdateBank(c *gin.Context) {
	var id models.Bank
	c.ShouldBindJSON(&id)
	tx, txerr := pg_db.Begin()
	if txerr != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx.Rollback()
		return
	}
	err := services.Update(tx,pg_db, id)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "Details could not be updated", "error": err})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "Details updated successfully"})
	}
}

func BranchDetails(c *gin.Context) {
	ans, err := services.Get_Branch(pg_db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of branch couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}
}
func InsertDetailsBranch(c *gin.Context) {
	var ans []models.Branch
	c.ShouldBindJSON(&ans)
	err := services.Insert_Branch(ans, pg_db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry data could not be entered in branch table"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data was inserted successfully"})
	}
}
func DeleteBranch(c *gin.Context) {
	var id models.Branch
	c.ShouldBindJSON(&id)
	tx, txerr := pg_db.Begin()
	if txerr != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx.Rollback()
		return
	}
	err := services.Delete_Branch(tx,pg_db, id)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry branch could not be deleted", "error": err})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "branch deleted successfully"})
	}
}
func UpdateBranch(c *gin.Context) {
	var id models.Branch
	c.ShouldBindJSON(&id)
	
	tx, txerr := pg_db.Begin()
	if txerr != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx.Rollback()
		return
	}
	err := services.Update_Branch(tx,pg_db, id)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "Details could not be updated", "error": err})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "Details updated successfully"})
	}
}

func GetAllBranches(c *gin.Context) {
	idStr := c.Param("id")

	// Convert the id string to uint
	id, _ := strconv.ParseUint(idStr, 10, 64)
	uintID := uint(id)
	ans, err := services.Get_All_Branches_By_Bankid(pg_db, uintID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of all branches couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}

}

func CustDetails(c *gin.Context) {
	ans, err := services.Details_Cust(pg_db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of customer couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}
}

// func insertDetailsCust(c *gin.Context) {
// 	var ans []models.Customer
// 	c.ShouldBindJSON(&ans)
// 	err := handler.Insertcust(ans, pg_db)

//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"message": "sorry data could not be entered in customers table"})
//		} else {
//			c.JSON(http.StatusOK, gin.H{"message": "Data was inserted successfully"})
//		}
//	}
func DeleteCust(c *gin.Context) {
	var id models.Customer
	c.ShouldBindJSON(&id)
	tx, txerr := pg_db.Begin()
	if txerr != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx.Rollback()
		return
	}
	err1 := services.Delete_Cust(tx,pg_db, id)
	if err1 != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry customer could not be deleted", "error": err1})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "customer deleted successfully"})
	}
}
func UpdateCust(c *gin.Context) {
	var id models.Customer
	c.ShouldBindJSON(&id)
	tx, txerr := pg_db.Begin()
	if txerr != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx.Rollback()
		return
	}
	err := services.Update_Cust(tx,pg_db, id)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "Details could not be updated", "error": err})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "Details updated successfully"})
	}
}
func GetAllCust(c *gin.Context) {
	idStr := c.Param("id")

	// Convert the id string to uint
	id, _ := strconv.ParseUint(idStr, 10, 64)
	uintID := uint(id)
	ans, err := services.Get_All_Customers_By_Branchid(pg_db, uintID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of all customers couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}

}

func OpenAcc(c *gin.Context) {
	var helper1 models.Openacc
	c.ShouldBindJSON(&helper1)

	var helper2 models.Account
	var helper3 []models.Customer
	helper2.Branch_id = helper1.Branch_id
	helper2.Balance = helper1.Balance
	helper2.Acc_type = helper1.Acc_type
	helper3 = append(helper3, helper1.Customer...)

	tx, txerr := pg_db.Begin()
	if txerr != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx.Rollback()
		return
	}

	err1 := services.Insert_Acc(tx, &helper2, pg_db)
	err2 := services.Insert_Cust(tx, &helper3, pg_db)

	if err1 != nil || err2 != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "sorry details of account and customer couldnt be inserted"})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "Data was inserted successfully"})
	}
	length := len(helper3)
	helper4 := make([]models.CustomerAccount, length)

	for i := range helper3 {
		helper4[i].AccountID = helper2.ID
		helper4[i].CustomerID = helper3[i].ID
	}
	tx1, txerr1 := pg_db.Begin()
	if txerr1 != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx1.Rollback()
		return
	}
	err := services.Insert_Ca(tx1, helper4)
	if err != nil {
		tx1.Rollback()
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "sorry customer account table couldnt be updated"})
	} else {
		tx1.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "Table was updated successfully"})
	}

}

func AccDetails(c *gin.Context) {
	ans, err := services.Details_Acc(pg_db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of account couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}
}

// func insertDetailsAcc(c *gin.Context) {
// 	var ans []models.Account
// 	c.ShouldBindJSON(&ans)
// 	err := handler.Insertacc(ans, pg_db)

//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"message": "sorry data could not be entered in accounts table"})
//		} else {
//			c.JSON(http.StatusOK, gin.H{"message": "Data was inserted successfully"})
//		}
//	}
func DeleteAcc(c *gin.Context) {
	var id models.Account
	c.ShouldBindJSON(&id)
	tx, txerr := pg_db.Begin()
	if txerr != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx.Rollback()
		return
	}
	err1 := services.Delete_Acc(tx, pg_db, id)
	if err1 != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry account could not be deleted", "error": err1})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "account deleted successfully"})
	}
}
func UpdateAcc(c *gin.Context) {
	var id models.Account
	c.ShouldBindJSON(&id)
	tx, txerr := pg_db.Begin()
	if txerr != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx.Rollback()
		return
	}
	err := services.Update_Acc(tx, pg_db, id)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "Details could not be updated", "error": err})

	} else {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "Details updated successfully"})

	}
}

func GetAllAcc(c *gin.Context) {
	idStr := c.Param("id")

	// Convert the id string to uint
	id, _ := strconv.ParseUint(idStr, 10, 64)
	uintID := uint(id)
	ans, err := services.Get_All_Accounts_By_Branchid(pg_db, uintID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of all accounts couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}

}

func TranDetails(c *gin.Context) {
	ans, err := services.Details_Tran(pg_db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of transaction couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}
}
func InsertDetailsTran(c *gin.Context) {
	var ans models.Transaction
	c.ShouldBindJSON(&ans)
	err := services.Insert_Tran(ans, pg_db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry data could not be entered in transactions table"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data was inserted successfully"})
	}
}
func DeleteTran(c *gin.Context) {
	var id models.Transaction
	c.ShouldBindJSON(&id)
	tx, txerr := pg_db.Begin()
	if txerr != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx.Rollback()
		return
	}

	err1 := services.Delete_Tran(tx,pg_db, id)
	if err1 != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry transaction could not be deleted", "error": err1})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "transaction deleted successfully"})
	}
}
func UpdateTran(c *gin.Context) {
	var id models.Transaction
	c.ShouldBindJSON(&id)
	tx, txerr := pg_db.Begin()
	if txerr != nil {
		log.Printf("error while starting transaction, Reason: %v\n", txerr)
		tx.Rollback()
		return
	}
	err := services.Update_Tran(tx,pg_db, id)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"message": "Details could not be updated", "error": err})
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, gin.H{"message": "Details updated successfully"})
	}
}

func GetAllTran(c *gin.Context) {
	idStr := c.Param("id")

	// Convert the id string to uint
	id, _ := strconv.ParseUint(idStr, 10, 64)
	uintID := uint(id)
	ans, err := services.Get_All_Transactions_By_Accountid(pg_db, uintID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of all transactions couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}

}

func CaDetailsById(c *gin.Context) {
	idStr := c.Param("id")

	// Convert the id string to uint
	id, _ := strconv.ParseUint(idStr, 10, 64)
	uintID := uint(id)
	ans, err := services.Get_Cust_By_Accid(pg_db, uintID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of all customers couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}

}
func CaDetailsByCid(c *gin.Context) {
	idStr := c.Param("id")

	// Convert the id string to uint
	id, _ := strconv.ParseUint(idStr, 10, 64)
	uintID := uint(id)
	ans, err := services.Get_Acc_By_Cid(pg_db, uintID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of accounts couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}

}