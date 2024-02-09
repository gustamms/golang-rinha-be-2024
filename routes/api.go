package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)

	transactionController := controllers.NewTransactionController()
	facades.Route().Post("/clientes/{id}/transacoes", transactionController.Create)

	extractController := controllers.NewExtractController()
	facades.Route().Get("/clientes/{id}/extrato", extractController.GetByConsumerId)
}
