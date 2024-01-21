<h1 align="center">
Go: Orientação a Objetos
</h1>

### Índice

* [Minha primeira struct](#Minha-primeira-struct)
* [Referência, ponteiro e métodos](#Referencia-ponteiro-metodos)
* [Retornos, pacotes e visibilidade](#Retornos-pacotes-visibilidade)
* [Composição e encapsulamento](#Composicao-encapsulamento)
* [Interface e novo tipo de conta](#Interface-novo-tipo-de-conta)



<h2 align="center" href="#Minha-primeira-struct">
Minha primeira struct
</h2>

* Criamos nossa primeira struct chamada ContaCorrente;

* Em seguida criamos um tipo com base na struct Conta corrente sem atribuir valor para os campos e descobrimos o Zero-initialization ou Inicialização zero;

* Para finalizar, vimos duas formas utilizar a struct criada: informando ou não os nomes dos campos.

<h2 align="center" href="#Referencia-ponteiro-metodos">
Referência, ponteiro e métodos
</h2>

* Criamos um nova conta corrente utilizando a palavra new;

* Em seguida, comparamos os tipos criados comparando suas referências e entendemos o que são ponteiros na prática;

* Para finalizar, desenvolvemos o método sacar que verifica se o valor do saque é maior do que zero e se a conta possui saldo.

<h2 align="center" href="#Retornos-pacotes-visibilidade">
Retornos, pacotes e visibilidade
</h2>

* Criamos um método chamado depositar, que quando invocado nos retorna uma mensagem (string) e o valor do novo saldo (float64);

* Em seguida, criamos o método transferir, que tira um valor de uma conta e transfere para uma conta destino reutilizando o método depositar;

* Para finalizar, criamos um pacote chamado contas e criamos um arquivo chamado contaCorrente.go para manter todo código referente a este tipo de conta.

<h2 align="center" href="#Composicao-encapsulamento">
Composição e encapsulamento
</h2>

* Criamos um novo pacote chamado clientes e um arquivo chamado clientes.go, onde desenvolvemos uma nova struct chamada Titular com nome, CPF e profissão;

* Em seguida, alteramos o campo titular da struct contaCorrente para ser do tipo da struct Titular;

* Para finalizar, alteramos a visibilidade do campo saldo e criamos um método chamado obterSaldo.

<h2 align="center" href="#Interface-novo-tipo-de-conta">
Interface e novo tipo de conta
</h2>

* Criamos um novo tipo de conta: a ContaPoupança;

* Para finalizar, criamos um novo tipo interface onde podemos utilizar tanto a conta corrente como poupança para pagar um boleto através da função Sacar.

### 🛠 Tecnologia

- [GoLang 1.20](https://go.dev/)


### 🛠  Ferramentas:

- [Compilador Online](https://go.dev/play/p/gkwKo7rholt)

