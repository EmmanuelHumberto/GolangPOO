package main

import (
	"banco/cliente"
	c "banco/contas"
	"fmt"
)

func main() {
	//contaCorrente01 := c.ContaCorrente{Titular: cliente.Titular{
	//	Nome:      "Juca Pirama",
	//	Cpf:       "74125896323",
	//	Profissão: "Desenvolvedor",
	//},
	//	NumeroAgencia: 01,
	//	NumeroConta:   12123,
	//	Saldo:         500}
	//
	//contaCorrente02 := c.ContaCorrente{Titular: cliente.Titular{
	//	Nome:      "SaciPerere",
	//	Cpf:       "74125896324",
	//	Profissão: "Desenvolvedor",
	//},
	//	NumeroAgencia: 01,
	//	NumeroConta:   12124,
	//	Saldo:         0}

	cliente01 := cliente.Titular{
		Nome:      "Juca Pirama",
		Cpf:       "74125896323",
		Profissão: "Desenvolvedor",
	}
	contaCorrente01 := c.ContaCorrente{
		Titular:       cliente01,
		NumeroAgencia: 01,
		NumeroConta:   12124,
		Saldo:         0,
	}
	cliente02 := cliente.Titular{
		Nome:      "Saci Perere",
		Cpf:       "74125896347",
		Profissão: "Desenvolvedor",
	}
	contaCorrente02 := c.ContaCorrente{
		Titular:       cliente02,
		NumeroAgencia: 01,
		NumeroConta:   12124,
		Saldo:         0,
	}

	contaCorrente01.Depositar(1000)
	contaCorrente01.Transferir(250, &contaCorrente02)
	fmt.Println(contaCorrente01.Saldo, contaCorrente02.Saldo)
}
