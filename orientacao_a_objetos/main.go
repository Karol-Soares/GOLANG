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

// definindo variaveis - aula 1
func main() {
	contaDoRoronoa := ContaCorrente{titular: "Roronoa",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5,
	}

	contaDaNami := ContaCorrente{"Nami", 222, 654321, 120}

	fmt.Println(contaDoRoronoa)
	fmt.Println(contaDaNami)
}
