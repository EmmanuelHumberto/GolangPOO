package main

import "fmt"

func main() {
	titular := "Emmanuel"
	numeroAgencia := 0001
	numeroConta := 12123
	saldo := float64(125.50)

	fmt.Println("Titular: ", titular,
		"\n Numero da agencia: ", numeroAgencia,
		"\n Numero da conta: ", numeroConta,
		"\n saldo R$: ", saldo)
}
