package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/pranav/my/repo/db"
	"github.com/pranav/my/repo/models"
	"github.com/pranav/my/repo/router"
)

var pg_db *pg.DB

func init() {
	orm.RegisterTable((*models.CustomerAccount)(nil))
}

func main() {

	app := gin.Default()
	pg_db = db.Connect()
	db.Create(pg_db)
	router.SetUpRoutes(app, pg_db)
	app.Run(":8080")

}


