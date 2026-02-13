package main

import (
	"fmt"
)

// primeiro exercicio multiplicar dois numeros inteiros
func main() {
	a := 3
	b := 7
	fmt.Println(a * b)

	// segundo exercicio somar dois numeros
	c := 2.5
	d := 4.7
	fmt.Println(d + c)

	// terceiro exercicio subtrair dois numeros
	e := 4
	f := 6
	fmt.Println(f - e)

	// quarto exercicio calculo de divisao
	g := 2
	h := 2
	fmt.Println(g / h)
	//exercicio definindo um numero par ou impar
	var numero int

	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	if numero%2 == 0 {
		fmt.Println("O número é PAR")
	} else {
		fmt.Println("O número é ÍMPAR")
	}
	//reconhecendo o nome de uma pessoa
	var nome string
	fmt.Print("Digite nome: ")
	fmt.Scan(&nome)
	if nome == "john" {
		fmt.Println("o nome é massa  ")
	} else {
		fmt.Println("é muito diferente ")

	}

}
