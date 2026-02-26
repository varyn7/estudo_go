package main

import "fmt"

func f1() { //não recebe nada como entrada de dados e não retrona nada
	fmt.Println("funbasicas")
}

func f2(p1 string, p2 string) { // recebe parametros e não retorna nada
	fmt.Printf("%s %s\n", p1, p2)
}

func f3() string { //não recebe parametro mas retorna algo no terminal
	return "f3"
}
func f4(p1, p2 string) string { //recebe dois parametro e retorna um unico parametro
	return fmt.Sprintf("f4:%s %s", p1, p2)

}
func f5() (string, string) {
	return "r1", "r2"
}
func main() {
	f1()
	f2("parametro1", "parametro2")
	r3, r4 := f3(), f4("parametro1", "parametro2")
	fmt.Println(r3, r4)
	r51, r52 := f5()
	fmt.Println("f5:", r51, r52)

}
