package routes

import (
	"github.com/labstack/echo/v4"
	"mytoko/controllers"
)

func AdminPath(e *echo.Echo, adminController *controllers.AdminController) {
	// ------------------------------------------------------------------
	// CRUD Admin
	// ------------------------------------------------------------------
	e.POST("/admin", adminController.CreateAdminController)
	e.GET("/admin", adminController.GetAllAdminController)
	e.DELETE("admin/:id", adminController.DeleteAdminController)
}