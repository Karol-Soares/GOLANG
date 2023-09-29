package main

import (
	"fmt"
)

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

// Criando metodo de transferencia entre contas

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}

}

func main() {
	contaDoRoronoa := ContaCorrente{titular: "Roronoa", saldo: 125.50}
	contaDoLuffy := ContaCorrente{titular: "Luffy", saldo: 200.0}

	status := contaDoRoronoa.Transferir(100, &contaDoLuffy)
	fmt.Println(status)

	fmt.Println(contaDoRoronoa)
	fmt.Println(contaDoLuffy)

}
