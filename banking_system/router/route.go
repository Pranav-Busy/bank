package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/pranav/my/repo/handler"
)

var pg_db *pg.DB

func SetUpRoutes(app *gin.Engine, db *pg.DB) {
	//bank end points
	pg_db = db
	bank := app.Group("/bank")
	bank.POST("/insert_details", handler.InsertDetails)
	bank.GET("/get_details", handler.BankDetails)
	bank.DELETE("/delete_bank", handler.DeleteBank)
	bank.PUT("/update_bank_details", handler.UpdateBank)

	//branch endpoints
	branch := app.Group("/branch")
	branch.GET("/branch_details", handler.BranchDetails)
	branch.POST("/insert_details_branch", handler.InsertDetailsBranch)
	branch.DELETE("/delete_branch", handler.DeleteBranch)
	branch.PUT("/update_branch", handler.UpdateBranch)
	branch.GET("/get_all_branch/:id", handler.GetAllBranches)
	//Customer endpoints
	cust := app.Group("/cust")
	cust.GET("/cust_details", handler.CustDetails)
	// cust.POST("/insert_details_cust", insertDetailsCust)
	cust.DELETE("/delete_cust", handler.DeleteCust)
	cust.PUT("/update_cust", handler.UpdateCust)
	cust.GET("/get_all_cust/:id", handler.GetAllCust)
	// Account endpoints
	acc := app.Group("/acc")
	acc.GET("/acc_details", handler.AccDetails)
	acc.POST("/open_acc", handler.OpenAcc)
	// acc.POST("/insert_details_acc", insertDetailsAcc)
	acc.DELETE("/delete_acc", handler.DeleteAcc)
	acc.PUT("/update_acc", handler.UpdateAcc)
	acc.GET("/get_all_acc/:id", handler.GetAllAcc)
	// transactions endpoints
	tran := app.Group("/tran")
	tran.GET("/tran_details", handler.TranDetails)
	tran.POST("/insert_details_tran", handler.InsertDetailsTran)
	tran.DELETE("/delete_tran", handler.DeleteTran)
	tran.PUT("/update_tran", handler.UpdateTran)
	tran.GET("/get_all_tran/:id", handler.GetAllTran)
	// Customer_account endpoints
	ca := app.Group("/ca")
	ca.GET("/ca_details/:id", handler.CaDetailsById)
	ca.GET("/acc_details/:id", handler.CaDetailsByCid)
}

func Returndb()(*pg.DB){
	return pg_db
}
