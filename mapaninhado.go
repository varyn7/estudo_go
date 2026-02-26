package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"g": {
			"gulano da silva": 1240,
			"guga da chagas":  1450,
		},
		"J": {
			"joao minino": 1800,
			"joao garoto": 2000,
		},
		"p": {
			"pedro paulo":  4000,
			"pedro poncio": 8000,
		},
	}
	delete(funcsPorLetra, "p")

	for letra, grupo := range funcsPorLetra {

		fmt.Println(letra, grupo)
	}
}
