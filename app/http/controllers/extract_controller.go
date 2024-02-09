package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type ExtractController struct {
	//Dependent services
}

func NewExtractController() *ExtractController {
	return &ExtractController{
		//Inject services
	}
}

func (r *ExtractController) GetByConsumerId(ctx http.Context) http.Response {
	id := ctx.Request().Input("id")

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": id,
	})
}
