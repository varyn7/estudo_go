package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} //quem conta esses numeros vai ser o compilador
	for i, numero := range numeros {
		fmt.Println(i, numero)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}
}
