package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tegarap/mytoko/config"
	"github.com/tegarap/mytoko/controllers"
	"github.com/tegarap/mytoko/lib/database"
	m "github.com/tegarap/mytoko/lib/middlewares"
	"github.com/tegarap/mytoko/routes"
)

func main()  {
	//db initialize database connection
	db := config.DBConnect()

	//initiate product model
	adminModel := database.NewAdminModel(db)
	cashierModel := database.NewCashierModel(db)
	productModel := database.NewProductModel(db)

	//initiate product controller
	newAdminController := controllers.NewAdminController(adminModel)
	newCashierController := controllers.NewCashierController(cashierModel)
	newProductController := controllers.NewControllerProduct(productModel)

	//create echo http
	e := echo.New()

	//register API path and controller
	routes.IndexPath(e)
	routes.AdminPath(e, newAdminController)
	routes.CashierPath(e, newCashierController)
	routes.ProductPath(e, newProductController)

	// run server
	address := fmt.Sprintf("localhost:%d", 8080)
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(address))
}
