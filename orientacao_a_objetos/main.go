package main

import "fmt"

// Banco

// Aplicação para gerenciar contas correntes
// Quais campos temos nas contas correntes?
// Titular: nome, Agencia: numero, Conta: numero, Saldo: valor
// queremos armazenar esses valores na memoria do computador

// Struct - aula 2

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// Criando metodo sacar
func (conta *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= conta.saldo
	if podeSacar {
		conta.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func main() {
	contaDoRoronoa := ContaCorrente{}
	contaDoRoronoa.titular = "Roronoa"
	contaDoRoronoa.saldo = 500

	fmt.Println(contaDoRoronoa.saldo)

	fmt.Println(contaDoRoronoa.Sacar(400))

	fmt.Println(contaDoRoronoa.saldo)
}
