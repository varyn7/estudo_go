package main

import "fmt"

// Declarar um numero e reconhecer se e par ou nao.
func main() {
	var n int
	fmt.Print("Digite um numero inteiro: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Entrada invalida.")
		return
	}
	if n%2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Impar")
	}

	// Exercicio de reconhecimento da maioridade.
	var idade int
	fmt.Print("Digite sua idade: ")
	if _, err := fmt.Scan(&idade); err != nil {
		fmt.Println("Entrada invalida.")
		return
	}
	if idade >= 18 {
		fmt.Println("Maior de idade")
	} else {
		fmt.Println("Menor de idade")
	}
	// Reconhecendo a positividade de um numero ou sua negatividade.
	var p int
	fmt.Print("Digite um numero inteiro: ")
	if _, err := fmt.Scan(&p); err != nil {
		fmt.Println("Entrada invalida.")
		return
	}
	if p > 0 {
		fmt.Println("Positivo")
	} else if p < 0 {
		fmt.Println("Negativo")
	} else {
		fmt.Println("Zero")
	}
}
