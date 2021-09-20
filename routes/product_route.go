package routes

import (
	"github.com/labstack/echo/v4"
	"mytoko/controllers"
	"net/http"
)

func IndexPath(e *echo.Echo)  {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "My Toko API")
	})
}

func ProductPath(e *echo.Echo, productController *controllers.ProductController) {
	// ------------------------------------------------------------------
	// CRUD Product
	// ------------------------------------------------------------------
	e.POST("/products", productController.CreateProductController)
	e.GET("/products", productController.GetAllProductController)
	e.GET("/products/:id", productController.GetProductByIdController)
	//e.PUT("/products/:id", customerController.EditCustomerController)
	//e.DELETE("/products/:id", customerController.DeleteCustomerController)
}