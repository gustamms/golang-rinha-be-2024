package controllers

import (
	"time"

	"github.com/goravel/framework/contracts/http"

	clientService "goravel/app/services"
	transactionService "goravel/app/services"
)

type ExtractController struct {
	//Dependent services
	clientService.ClientService
	transactionService.TransactionService
}

type Response struct {
	Balance
	LastTransactions
}

type Balance struct {
	Total       int
	DataExtrato string
	Limite      int
}

type LastTransactions struct {
	Transactions
}

type Transactions struct {
	Valor       int
	Tipo        string
	Descricao   string
	RealizadaEm string
}

func NewExtractController() *ExtractController {
	return &ExtractController{
		//Inject services
	}
}

func (r *ExtractController) GetByConsumerId(ctx http.Context) http.Response {
	id := ctx.Request().Input("id")

	client := clientService.NewClientService()
	transactionService := transactionService.NewTransactionService()

	clientInformations := client.GetById(id)

	if clientInformations.Id == 0 {
		return ctx.Response().Json(http.StatusUnprocessableEntity, http.Json{
			"message": "Usuário não existente",
		})
	}

	transactions := transactionService.GetLastTransactions(clientInformations.Id)

	balance := &Balance{
		Total:       clientInformations.Balance,
		DataExtrato: time.Now().UTC().String(),
		Limite:      clientInformations.Limit_account,
	}
	// ultimasTransacoes := []Transactions{
	// 	Transactions{
	// 		Valor: 1,
	// 	},
	// }

	return ctx.Response().Json(http.StatusOK, http.Json{
		"saldo":              balance,
		"ultimas_transacoes": transactions,
	})
}
