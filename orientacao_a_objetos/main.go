package main

import (
	"fmt"
	"orientacao_a_objetos/contas"
)

func main() {
	contaDoRoronoa := contas.ContaCorrente{titular: "Roronoa", Saldo: 125.50}
	contaDoLuffy := contas.ContaCorrente{titular: "Luffy", Saldo: 200.0}

	status := contaDoRoronoa.Transferir(100, &contaDoLuffy)
	fmt.Println(status)

	fmt.Println(contaDoRoronoa)
	fmt.Println(contaDoLuffy)

}
