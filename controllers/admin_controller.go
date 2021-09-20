package controllers

import (
	"github.com/labstack/echo/v4"
	"mytoko/lib/database"
	"mytoko/models"
	"net/http"
	"strconv"
)

type AdminController struct {
	adminModel database.AdminModel
}

func NewAdminController(adminModel database.AdminModel) *AdminController {
	return &AdminController{
		adminModel,
	}
}

func (controller *AdminController) CreateAdminController(c echo.Context) error {
	var adminRequest models.Admin

	if err := c.Bind(&adminRequest); err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	admin := models.Admin{
		Username: adminRequest.Username,
		Password: adminRequest.Password,
	}

	if err := controller.adminModel.Add(admin); err != nil {
		return c.JSON(http.StatusInternalServerError, NewInternalServerErrorResponse())
	}

	return c.JSON(http.StatusOK, NewSuccessOperationResponse())
}

func (controller *AdminController) GetAllAdminController(c echo.Context) error {
	admins, err := controller.adminModel.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, admins)
}

func (controller *AdminController) DeleteAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	product, err := controller.adminModel.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, product)
}