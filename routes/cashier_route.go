package routes

import (
	"github.com/labstack/echo/v4"
	"mytoko/controllers"
)

func CashierPath(e *echo.Echo, cashierController *controllers.CashierController) {
	// ------------------------------------------------------------------
	// CRUD Cashier
	// ------------------------------------------------------------------
	e.POST("/cashiers", cashierController.CreateCashierController)
	e.GET("/cashiers", cashierController.GetAllCashierController)
	e.GET("/cashiers/:id", cashierController.GetCashierByIdController)
	//e.PUT("/cashiers/:id", customerController.EditCustomerController)
	//e.DELETE("/cashiers/:id", customerController.DeleteCustomerController)
}