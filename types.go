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

}
