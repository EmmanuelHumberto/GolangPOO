package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaCorrente01 := ContaCorrente{titular: "Emmanuel", numeroAgencia: 01, numeroConta: 12123, saldo: 2000}
	contaCorrente02 := ContaCorrente{"Saci Perere", 01, 12133, 4000}
	fmt.Println(contaCorrente01, contaCorrente02)
}
