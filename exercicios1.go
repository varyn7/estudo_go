package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 42
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))

	// Exemplo de tipo definido e conversao para o tipo subjacente
	type MeuInt int
	var z MeuInt = 7
	fmt.Println(z)
	fmt.Println(reflect.TypeOf(z))

	uz := int(z)
	fmt.Println(uz)
	fmt.Println(reflect.TypeOf(uz))

	var y float64 = 4.2
	fmt.Println(y)
	fmt.Println(reflect.TypeOf(y))

}
