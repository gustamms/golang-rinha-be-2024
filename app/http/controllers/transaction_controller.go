package controllers

import (
	transactionService "goravel/app/services"

	"github.com/goravel/framework/contracts/http"
)

type TransactionController struct {
	//Dependent services
	transactionService.TransactionService
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		//Inject services
	}
}

func (r *TransactionController) Create(ctx http.Context) http.Response {
	validator, err := ctx.Request().Validate(map[string]string{
		"valor":     "required",
		"tipo":      "required|string|in:c,d",
		"descricao": "required|max_len:10|string",
	})

	id := ctx.Request().Input("id")
	value := ctx.Request().Input("valor")
	tipo := ctx.Request().Input("tipo")
	description := ctx.Request().Input("descricao")

	if err != nil {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"message": err.Error(),
		})
	}
	if validator.Fails() {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"message": validator.Errors().All(),
		})
	}

	service := transactionService.NewTransactionService()
	response, err := service.Create(id, tipo, description, value)

	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": err,
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"transaction": response,
		"message":     "Criado com sucesso!",
	})
}
