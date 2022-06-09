package contas

import "banco/cliente"

type ContaCorrente struct {
	//letas minisculas: private
	//letas maiusculas: public
	Titular       cliente.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDosaque float64) (string, float64) {
	podeSacar := valorDosaque <= c.saldo && valorDosaque > 0
	if podeSacar {
		c.saldo -= valorDosaque
		return "Saque realizado!", c.saldo
	} else {
		return "Saldo insuficiente", c.saldo
	}
}
func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.saldo += valorDeposito
		return "Deposito realizado!", c.saldo
	} else {
		return "Deposito de valor negativo não pode ser reaizado", c.saldo
	}
}
func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	podetransferir := valorTransferencia > 0 && valorTransferencia <= c.saldo
	if podetransferir {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false

	}
}
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo

}
