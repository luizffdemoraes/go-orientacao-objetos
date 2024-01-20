package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// conta := ContaCorrente{
	// 	titular:       "Luiz",
	// 	numeroAgencia: 589,
	// 	numeroConta:   123456,
	// 	saldo:         125.50,
	// }

	// conta2 := ContaCorrente{
	// 	titular:       "Luiz",
	// 	numeroAgencia: 589,
	// 	numeroConta:   123456,
	// 	saldo:         125.50,
	// }

	// Validar igualdade
	// fmt.Println(conta == conta2)

	// contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	// fmt.Println(conta)
	// fmt.Println(contaDaBruna)

	// Ponteiro *
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500
	contaDaCris.numeroConta = 111333
	contaDaCris.numeroAgencia = 300

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"
	contaDaCris2.saldo = 500
	contaDaCris2.numeroConta = 111333
	contaDaCris2.numeroAgencia = 300

	// Visualizar o endereço
	fmt.Println(&contaDaCris)

	// Valída o endereço
	fmt.Println(contaDaCris == contaDaCris2)
	// Valída o conteúdo
	fmt.Println(*contaDaCris == *contaDaCris2)
}
