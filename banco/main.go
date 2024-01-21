package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Silvia", CPF: "123.111.123.12", Profissao: "Desenvolvedor"}, NumeroConta: 123, NumeroAgencia: 123456, Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Gustavo", CPF: "123.111.123.13", Profissao: "Tech Leader"}, NumeroConta: 123, NumeroAgencia: 123456, Saldo: 300}

	// status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	status := contaDoGustavo.Transferir(-200, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
