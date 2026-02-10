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
	//ponto literal com numero flutuante retorna sempre como float64

	var x float64 = 3.14
	fmt.Println("o tipo de x e", reflect.TypeOf(x))
	fmt.Println("o tipo literal de 3.14 e", reflect.TypeOf(3.14))

	//boleanos
	bol1 := true
	fmt.Println("o tipo de bol1 e", reflect.TypeOf(bol1))
	bol2 := false
	fmt.Println("o tipo de bol2 e", reflect.TypeOf(bol2))
	//string
	s1 := "ola"
	s2 := "john"
	fmt.Println("o tipo de s1 e", reflect.TypeOf(s1))
	fmt.Println("o tipo de s2 e", reflect.TypeOf(s2))
	r2 := []rune(s1)[0]
	fmt.Println("o tipo de r2 e", reflect.TypeOf(r2))
	//funcao len
	fmt.Println("o respectivo tamanho da string e", len(s1+s2))

	//string de multiplas linhas
	s3 := `ola
meu
nome
e
john`
	fmt.Println("o tamanho da string e", len(s3))
}
