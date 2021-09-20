package controllers

import (
	"github.com/labstack/echo/v4"
	"mytoko/lib/database"
	"mytoko/models"
	"net/http"
	"strconv"
)

type ProductController struct {
	productModel database.ProductModel
}

func NewControllerProduct(productModel database.ProductModel) *ProductController {
	return &ProductController{
		productModel,
	}
}

func (controller *ProductController) CreateProductController(c echo.Context) error {
	var productRequest models.Product

	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	product := models.Product{
		Sku:   productRequest.Sku,
		Name:  productRequest.Name,
	}

	err := controller.productModel.Add(product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewInternalServerErrorResponse())
	}

	return c.JSON(http.StatusOK, NewSuccessOperationResponse())
}

func (controller *ProductController) GetAllProductController(c echo.Context) error {
	products, err := controller.productModel.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, products)
}

func (controller *ProductController) GetProductByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	product, err := controller.productModel.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, product)
}