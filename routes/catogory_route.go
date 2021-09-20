package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/mytoko/controllers"
)

func CategoryPath(e *echo.Echo, categoryController *controllers.CategoryController) {
	// ------------------------------------------------------------------
	// CRUD Categories
	// ------------------------------------------------------------------
	e.POST("/categories", categoryController.CreateCategoryController)
	e.GET("/categories", categoryController.GetAllCategoryController)
	e.PUT("/categories/:id", categoryController.UpdateCategoryController)
	e.DELETE("/categories/:id", categoryController.DeleteCategoryController)
}