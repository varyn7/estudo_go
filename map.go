package main

import "fmt"

func main() {
	// mapas devem ser inicializados
	aprovado := make(map[int64]string)
	aprovado[12345678] = "joao jaime"
	aprovado[95665755] = "neosanildo"
	aprovado[7867965984] = "joao minino"
	fmt.Println(aprovado)
	for cpf, nome := range aprovado {
		fmt.Printf("%d - %s \n", cpf, nome)

	}
	fmt.Println(aprovado[12345678])
	delete(aprovado, 12345678)
	fmt.Println(aprovado[12345678])
}
