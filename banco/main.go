package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDaLuisa := contas.ContaPoupanca{}
	contaDaLuisa.Depositar(500)
	// Passa o endereço da conta
	PagarBoleto(&contaDaLuisa, 100)

	fmt.Println(contaDaLuisa.ObterSaldo())
}
