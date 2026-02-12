package main

import "fmt"

func main() {
	var a byte = 3

	fmt.Println(a)
	i := 3 //inferencia de tipo
	i += 3 // inferencia de adição
	i -= 3 // inferencia subtrativa
	i *= 3 //inferencia multiplicativa
	i /= 3 //inferencia divisa
	i %= 3 // inferencia de modulo
	fmt.Println(i)
	//atribuição de valores e alternando a ordem
	x, y := 1, 2
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}
