package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(22, 44, 600)
	fmt.Println("o literal inteiro equivale a:", reflect.TypeOf(32000))
	a := 255
	fmt.Println("o byte equivale a:", reflect.TypeOf(a))
	b := math.MaxInt64
	fmt.Println("o literal inteiro maximo equivale a:", b)
	fmt.Println("o type de variavel do b e", reflect.TypeOf(b))
	var c byte = 140
	fmt.Println("o byte equivale c:", reflect.TypeOf(c))
	var r rune = rune(a)
	fmt.Println("o rune equivale:", reflect.TypeOf(r))
	fmt.Println(r)
	var x float64 = 3.14
	fmt.Println("o tipo de x e", reflect.TypeOf(x))
	fmt.Println("o tipo literal de 3.14 e", reflect.TypeOf(3.14))
}
