package contas

import "banco/cliente"

type ContaCorrente struct {
	//letas minisculas: private
	//letas maiusculas: public
	Titular       cliente.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDosaque float64) (string, float64) {
	podeSacar := valorDosaque <= c.Saldo && valorDosaque > 0
	if podeSacar {
		c.Saldo -= valorDosaque
		return "Saque realizado!", c.Saldo
	} else {
		return "Saldo insuficiente", c.Saldo
	}
}
func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.Saldo += valorDeposito
		return "Deposito realizado!", c.Saldo
	} else {
		return "Deposito de valor negativo não pode ser reaizado", c.Saldo
	}
}
func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	podetransferir := valorTransferencia > 0 && valorTransferencia <= c.Saldo
	if podetransferir {
		c.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false

	}
}
