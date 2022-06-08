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
	contaCorrente04 := ContaCorrente{titular: "Emmanuel", numeroAgencia: 01, numeroConta: 12123, saldo: 2000}
	contaCorrente02 := ContaCorrente{"Saci Perere", 01, 12133, 4000}
	fmt.Println(contaCorrente01, contaCorrente02)

	//comprando conetudo e endereços das structs
	fmt.Println(contaCorrente01 == contaCorrente04)

	contaCorrente03 := new(ContaCorrente)
	contaCorrente03.titular = "Juca Pirama"
	contaCorrente03.numeroAgencia = 01
	contaCorrente03.numeroConta = 10102
	contaCorrente03.saldo = 10000

	contaCorrente05 := new(ContaCorrente)
	contaCorrente05.titular = "Juca Pirama"
	contaCorrente05.numeroAgencia = 01
	contaCorrente05.numeroConta = 10102
	contaCorrente05.saldo = 10000

	//comprando conetudo e endereços das structs
	fmt.Println(contaCorrente03 == contaCorrente05)
	fmt.Println(*contaCorrente03 == *contaCorrente05)
	fmt.Println(&contaCorrente03, &contaCorrente05)

}
