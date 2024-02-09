package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type TransactionController struct {
	//Dependent services
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		//Inject services
	}
}

func (r *TransactionController) Create(ctx http.Context) http.Response {

	validator, err := ctx.Request().Validate(map[string]string{
		"valor":     "required|int|min_len:0",
		"tipo":      "required|string|in:c,d",
		"descricao": "required|max_len:10|string",
	})

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

	id := ctx.Request().Input("id")

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": id,
	})
}
