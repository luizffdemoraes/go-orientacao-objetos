package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	conta := ContaCorrente{
		titular:       "Luiz",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.50,
	}

    contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println(conta)
    fmt.Println(contaDaBruna)
}
