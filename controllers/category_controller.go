package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/mytoko/lib/database"
	"github.com/tegarap/mytoko/models"
	"net/http"
	"strconv"
)

type CategoryController struct {
	categoryModel database.CategoryModel
}

func NewCategoryModel(categoryModel database.CategoryModel) *CategoryController {
	return &CategoryController{
		categoryModel,
	}
}

func (controller *CategoryController) CreateCategoryController(c echo.Context) error {
	var categoryRequest models.Category

	if err := c.Bind(&categoryRequest); err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	category := models.Category{
		Name: categoryRequest.Name,
	}

	err := controller.categoryModel.Add(category)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewInternalServerErrorResponse())
	}

	return c.JSON(http.StatusOK, NewSuccessOperationResponse())
}

func (controller *CategoryController) GetAllCategoryController(c echo.Context) error {
	categories, err := controller.categoryModel.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, categories)
}

func (controller *CategoryController) UpdateCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	var categoryRequest models.Category

	if err = c.Bind(&categoryRequest); err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	category := models.Category{
		Name: categoryRequest.Name,
	}

	if _, err = controller.categoryModel.Update(category, id); err != nil {
		return c.JSON(http.StatusNotFound, NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, NewSuccessOperationResponse())
}

func (controller *CategoryController) DeleteCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	product, err := controller.categoryModel.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, product)
}