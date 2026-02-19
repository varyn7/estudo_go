package main

import "fmt"

func notaParaconceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "a"
	case 8, 7:
		return "b"
	case 6, 5:
		return "c"
	case 4, 3:
		return "d"
	default:
		return "nota anulada"
	}
}

func main() {
	// Exemplo simples para permitir executar o pacote main.
	fmt.Println(notaParaconceito(10))
	fmt.Println(notaParaconceito(9))
	fmt.Println(notaParaconceito(8))
	fmt.Println(notaParaconceito(7))
	fmt.Println(notaParaconceito(6))
	fmt.Println(notaParaconceito(7.6))

}
