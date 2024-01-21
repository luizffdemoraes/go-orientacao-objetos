package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Silvia"}, Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Gustavo"}, Saldo: 100}

	// status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	status := contaDoGustavo.Transferir(-200, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
