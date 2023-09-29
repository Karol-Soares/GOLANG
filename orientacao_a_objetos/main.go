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

// Criando metodo depositar com múltiplos retornos
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "deposito realizado com sucesso", c.saldo
	} else {
		return "valor do deposito menor que zero", c.saldo
	}
}

func main() {
	contaDoRoronoa := ContaCorrente{}
	contaDoRoronoa.titular = "Roronoa"
	contaDoRoronoa.saldo = 500

	fmt.Println(contaDoRoronoa.saldo)
	status, valor := contaDoRoronoa.Depositar(2000)
	fmt.Println(status, valor)

}
