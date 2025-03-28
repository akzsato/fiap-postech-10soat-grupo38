package entities

import (
	"errors"
	"strings"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   string
}

func ClienteNew(nome, email, CPF string) (*Cliente, error) {

	if  strings.TrimSpace(nome) == "" || strings.TrimSpace(email) == "" || strings.TrimSpace(CPF) == "" {
		return nil, errors.New("nenhum dos campos podem estar em branco")
	}

	return &Cliente{
		Nome:  nome,
		Email: email,
		CPF:   CPF,
	}, nil
}