package main

import (
	"fmt"

	"strconv"
	"strings"

	"github.com/fuzzyqu/simplified-des/sdes"
)

// letra mapeia 4 bits para uma letra.
func letra(bits string) string {
	switch bits {
	case "0000":
		return "A"
	case "0001":
		return "B"
	case "0010":
		return "C"
	case "0011":
		return "D"
	case "0100":
		return "E"
	case "0101":
		return "F"
	case "0110":
		return "G"
	case "0111":
		return "H"
	case "1000":
		return "I"
	case "1001":
		return "J"
	case "1010":
		return "K"
	case "1011":
		return "L"
	case "1100":
		return "M"
	case "1101":
		return "N"
	case "1110":
		return "O"
	case "1111":
		return "P"
	default:
		return "Z"
	}
}

func main() {

	texto := [8]uint8{1, 0, 1, 0, 0, 0, 1, 0}
	chave := [10]uint8{0, 1, 1, 1, 1, 1, 1, 1, 0, 1}

	decifrado := sdes.Descriptografar(texto, chave)

	fmt.Println("Original:\t", texto)
	fmt.Println("Decifrado:\t", decifrado)

	// Converte a array de uint8 para uma slice de string.
	decifradoSlice := make([]string, 0, 8)
	for _, v := range decifrado {
		decifradoSlice = append(decifradoSlice, strconv.Itoa(int(v)))
	}

	// Cria uma string com cada metade da array.
	e := strings.Join(decifradoSlice[:4], "")
	d := strings.Join(decifradoSlice[4:], "")

	// Imprime as letras correspondentes.
	fmt.Printf("Código convertido: '%s'\n", letra(e)+letra(d))
}
