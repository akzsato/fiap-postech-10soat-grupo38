package payment

import (
	"math/rand"
	"time"
)

type PaymentRequest struct {
	IdPagamento string
	Valor       string
	Status      string
	DataCriacao string
}

func ProcessPayment(request PaymentRequest) (string) {

	rand.Seed(time.Now().UnixNano())

	result := chooseWithProbability()

	return result
}

func chooseWithProbability() string {
	
	randomNumber := rand.Intn(100)

	if randomNumber < 90 {
		return "Recebido"
	}
	return "Negado"
}