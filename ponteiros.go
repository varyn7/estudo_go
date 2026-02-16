package main

import "fmt"

// golang nao tem operacoes aritmeticas com ponteiros
func main() {
	i := 1
	var p *int = nil
	// pegando endereço da variável
	p = &i
	*p++ // desreferenciando o ponteiro (pegando o valor )
	i++
	fmt.Println(*p, i, p)
}
