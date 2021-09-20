package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"mytoko/config"
	"mytoko/controllers"
	"mytoko/lib/database"
	m "mytoko/lib/middlewares"
	"mytoko/routes"
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
	address := fmt.Sprintf("localhost:%d", 5050)
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(address))
}