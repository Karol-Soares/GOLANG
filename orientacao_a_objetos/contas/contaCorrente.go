package contas

// Banco

// Aplicação para gerenciar contas correntes
// Quais campos temos nas contas correntes?
// Titular: nome, Agencia: numero, Conta: numero, Saldo: valor
// queremos armazenar esses valores na memoria do computador

// Struct - aula 2

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

// Criando metodo sacar
func (conta *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= conta.Saldo
	if podeSacar {
		conta.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

// Criando metodo depositar com múltiplos retornos
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "deposito realizado com sucesso", c.Saldo
	} else {
		return "valor do deposito menor que zero", c.Saldo
	}
}

// Criando metodo de transferencia entre contas

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}

}
