package main

import "fmt"

// Banco

// Aplicação para gerenciar contas correntes
// Quais campos temos nas contas correntes?
// Titular: nome, Agencia: numero, Conta: numero, Saldo: valor
// queremos armazenar esses valores na memoria do computador

func main() {
	var titular string = "Roronoa"
	var numeroAgencia int = 589
	var numeroConta int = 123456
	var saldo float64 = 125.5

	fmt.Println(titular, numeroAgencia, numeroConta, saldo)

	var titular2 string = "Luffy"
	var numeroAgencia2 int = 985
	var numeroConta2 int = 654321
	var saldo2 float64 = 100.5

	fmt.Println(titular2, numeroAgencia2, numeroConta2, saldo2)
}
