package main

import "fmt"

func main() {
	var nota float64
	fmt.Print("Digite a nota: ")
	fmt.Scanln(&nota)
	if nota >= 6 {
		fmt.Println(nota, "aprovado")
	} else {
		fmt.Println(nota, "reprovado")
	}
}
