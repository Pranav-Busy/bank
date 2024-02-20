package router

import (
	"net/http"
	
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/pranav/my/repo/handler"
	"github.com/pranav/my/repo/models"
)

var pg_db *pg.DB

func SetUpRoutes(app *gin.Engine, db *pg.DB) {
	//bank end points
	pg_db = db
	bank := app.Group("/bank")
	bank.POST("/insert_details", insertDetails)
	bank.GET("/get_details", bankDetails)
	bank.DELETE("/delete_bank", deleteBank)
	bank.PUT("/update_bank_details", updateBank)

	//branch endpoints
	branch := app.Group("/branch")
	branch.GET("/branch_details", branchDetails)
	branch.POST("/insert_details_branch", insertDetailsBranch)
	branch.DELETE("/delete_branch", deleteBranch)
	branch.PUT("/update_branch", updateBranch)
	branch.GET("/get_all_branch/:id",getAllBranches)
	//Customer endpoints
	cust := app.Group("/cust")
	cust.GET("/cust_details", custDetails)
	// cust.POST("/insert_details_cust", insertDetailsCust)
	cust.DELETE("/delete_cust", deleteCust)
	cust.PUT("/update_cust", updateCust)
	cust.GET("/get_all_cust/:id",getAllCust)
	// Account endpoints
	acc := app.Group("/acc")
	acc.GET("/acc_details", accDetails)
	acc.POST("/open_acc",openAcc)
	// acc.POST("/insert_details_acc", insertDetailsAcc)
	acc.DELETE("/delete_acc", deleteAcc)
	acc.PUT("/update_acc", updateACC)
    acc.GET("/get_all_acc/:id",getAllAcc)
	// transactions endpoints
	tran := app.Group("/tran")
	tran.GET("/tran_details", tranDetails)
	tran.POST("/insert_details_tran", insertDetailsTran)
	tran.DELETE("/delete_tran", deleteTran)
	tran.PUT("/update_tran", updateTran)
	tran.GET("/get_all_tran/:id",getAllTran)
	// Customer_account endpoints
	ca := app.Group("/ca")
	ca.GET("/ca_details/:id",caDetailsbyid)
	ca.GET("/acc_details/:id",caDetailsbycid)
}

func insertDetails(c *gin.Context) {
	var ans []models.Bank
	c.ShouldBindJSON(&ans)
	err := handler.Insert(ans, pg_db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry data could not be entered in bank table"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data was inserted successfully"})
	}
}

func bankDetails(c *gin.Context) {
	ans, err := handler.Details(pg_db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of bank couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}
}
func deleteBank(c *gin.Context) {
	var id models.Bank
	c.ShouldBindJSON(&id)
	err := handler.Delete(pg_db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry bank could not be deleted", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "bank deleted successfully"})
	}
}

func updateBank(c *gin.Context) {
	var id models.Bank
	c.ShouldBindJSON(&id)
	err := handler.Update(pg_db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Details could not be updated", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details updated successfully"})
	}
}

func branchDetails(c *gin.Context) {
	ans, err := handler.GetBranch(pg_db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of branch couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}
}
func insertDetailsBranch(c *gin.Context) {
	var ans []models.Branch
	c.ShouldBindJSON(&ans)
	err := handler.InsertBranch(ans, pg_db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry data could not be entered in branch table"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data was inserted successfully"})
	}
}
func deleteBranch(c *gin.Context) {
	var id models.Branch
	c.ShouldBindJSON(&id)
	err := handler.DeleteBranch(pg_db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry branch could not be deleted", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "branch deleted successfully"})
	}
}
func updateBranch(c *gin.Context) {
	var id models.Branch
	c.ShouldBindJSON(&id)
	err := handler.UpdateBranch(pg_db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Details could not be updated", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details updated successfully"})
	}
}

func getAllBranches( c *gin.Context){
	idStr := c.Param("id")

	// Convert the id string to uint
	id, _ := strconv.ParseUint(idStr, 10, 64)
    uintID := uint(id)
	ans,err:= handler.Getallbranchesbybankid(pg_db,uintID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of all branches couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}

}

func custDetails(c *gin.Context) {
	ans, err := handler.Detailscust(pg_db)
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

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry data could not be entered in customers table"})
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"message": "Data was inserted successfully"})
// 	}
// }
func deleteCust(c *gin.Context) {
	var id models.Customer
	c.ShouldBindJSON(&id)
	err1 := handler.Deletecust(pg_db, id)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry customer could not be deleted", "error": err1})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "customer deleted successfully"})
	}
}
func updateCust(c *gin.Context) {
	var id models.Customer
	c.ShouldBindJSON(&id)
	err := handler.Updatecust(pg_db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Details could not be updated", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details updated successfully"})
	}
}
func getAllCust( c *gin.Context){
	idStr := c.Param("id")

	// Convert the id string to uint
	id, _ := strconv.ParseUint(idStr, 10, 64)
    uintID := uint(id)
	ans,err:= handler.Getallcustomersbybranchid(pg_db,uintID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of all customers couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}

}

func openAcc(c *gin.Context){
	var helper1 models.Openacc
	 c.ShouldBindJSON(&helper1)
	var helper2  models.Account
	var helper3  []models.Customer
	helper2.Branch_id=helper1.Branch_id
	helper2.Balance=helper1.Balance
	helper2.Acc_type=helper1.Acc_type
	helper3 = append(helper3, helper1.Customer...)
	
	

	
	
	 err1:=handler.Insertacc(&helper2, pg_db)
	 err2:=handler.Insertcust(&helper3,pg_db)
	

	 if err1!=nil || err2!=nil{
		c.JSON(http.StatusBadRequest,
		  gin.H{"message": "sorry details of account and customer couldnt be inserted"})
	 }else 
	 {
		c.JSON(http.StatusOK, gin.H{"message": "Data was inserted successfully"})
	 }
	 length := len(helper3)
     helper4 := make([]models.CustomerAccount, length)

	
	 for i :=range helper3{
		helper4[i].AccountID=helper2.ID
		helper4[i].CustomerID=helper3[i].ID
	 }
	 err:=handler.Insertca(pg_db,helper4);
	 if err!=nil{
		c.JSON(http.StatusBadRequest,
		  gin.H{"message": "sorry customer account table couldnt be updated"})
	 }else 
	 {
		c.JSON(http.StatusOK, gin.H{"message": "Table was updated successfully"})
	 }

}

func accDetails(c *gin.Context) {
	ans, err := handler.Detailsacc(pg_db)
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

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry data could not be entered in accounts table"})
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"message": "Data was inserted successfully"})
// 	}
// }
func deleteAcc(c *gin.Context) {
	var id models.Account
	c.ShouldBindJSON(&id)
	err1 := handler.Deleteacc(pg_db, id)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry account could not be deleted", "error": err1})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "account deleted successfully"})
	}
}
func updateACC(c *gin.Context) {
	var id models.Account
	c.ShouldBindJSON(&id)
	err := handler.Updateacc(pg_db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Details could not be updated", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details updated successfully"})
	}
}

func getAllAcc( c *gin.Context){
	idStr := c.Param("id")

	// Convert the id string to uint
	id, _ := strconv.ParseUint(idStr, 10, 64)
    uintID := uint(id)
	ans,err:= handler.Getallaccountsbybranchid(pg_db,uintID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of all accounts couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}

}


func tranDetails(c *gin.Context) {
	ans, err := handler.Detailstran(pg_db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of transaction couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}
}
func insertDetailsTran(c *gin.Context) {
	var ans []models.Transaction
	c.ShouldBindJSON(&ans)
	err := handler.Inserttran(ans, pg_db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry data could not be entered in transactions table"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Data was inserted successfully"})
	}
}
func deleteTran(c *gin.Context) {
	var id models.Transaction
	c.ShouldBindJSON(&id)
	err1 := handler.Deletetran(pg_db, id)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry transaction could not be deleted", "error": err1})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "transaction deleted successfully"})
	}
}
func updateTran(c *gin.Context) {
	var id models.Transaction
	c.ShouldBindJSON(&id)
	err := handler.Updatetran(pg_db, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Details could not be updated", "error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details updated successfully"})
	}
}

func getAllTran( c *gin.Context){
	idStr := c.Param("id")

	// Convert the id string to uint
	id, _ := strconv.ParseUint(idStr, 10, 64)
    uintID := uint(id)
	ans,err:= handler.Getalltransactionsbyaccountid(pg_db,uintID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of all transactions couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}

}

func caDetailsbyid ( c *gin.Context){
	idStr := c.Param("id")

	// Convert the id string to uint
	id, _ := strconv.ParseUint(idStr, 10, 64)
    uintID := uint(id)
	ans,err:= handler.Getcustbyaccid(pg_db,uintID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of all customers couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}

}
func caDetailsbycid ( c *gin.Context){
	idStr := c.Param("id")

	// Convert the id string to uint
	id, _ := strconv.ParseUint(idStr, 10, 64)
    uintID := uint(id)
	ans,err:= handler.Getaccbycid(pg_db,uintID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry details of accounts couldnt be fetched"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Details fetched successfully", "Result": ans})
	}

}


