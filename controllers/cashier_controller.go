package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tegarap/mytoko/lib/database"
	"github.com/tegarap/mytoko/models"
	"net/http"
	"strconv"
)

type CashierController struct {
	cashierModel database.CashierModel
}

func NewCashierController(cashierModel database.CashierModel) *CashierController {
	return &CashierController{
		cashierModel,
	}
}

func (controller *CashierController) CreateCashierController(c echo.Context) error {
	var cashierRequest models.Cashier

	if err := c.Bind(&cashierRequest); err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	cashier := models.Cashier{
		Name: cashierRequest.Name,
		Username: cashierRequest.Username,
		Password: cashierRequest.Password,
	}

	err := controller.cashierModel.Add(cashier)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewInternalServerErrorResponse())
	}

	return c.JSON(http.StatusOK, NewSuccessOperationResponse())
}

func (controller *CashierController) GetAllCashierController(c echo.Context) error {
	cashiers, err := controller.cashierModel.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, cashiers)
}

func (controller *CashierController) GetCashierByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	cashier, err := controller.cashierModel.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, cashier)
}