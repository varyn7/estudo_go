package main

import (
	"fmt"
	"math"
)

// funcao de soma (adicao)
func main() {
	a := 2
	b := 2
	fmt.Println("o resultado dessa soma é", a+b)

	// operador de subtracao
	c := 2
	d := 2
	fmt.Println("o resultado dessa subtracao é", c-d)
	// operador de multiplicacao
	e := 2
	f := 2
	fmt.Println("o resultado dessa multiplicacao é", e*f)
	// operador de divisao
	g := 2
	h := 2
	fmt.Println("o resultado dessa divisao é", g/h)
	//operador de modulo
	i := 2
	j := 2
	fmt.Println("o resultado desse resto de divisao é", i%j)
	// operadores usando math...
	k := 2.5
	l := 3.5
	fmt.Println("o maior valor é", math.Max(k, l))
	fmt.Println("o menor valor dos dois é", math.Min(k, l))
	fmt.Println("a exponenciação é", math.Pow(k, l))
}
