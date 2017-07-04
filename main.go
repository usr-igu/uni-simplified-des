package main

import (
	"fmt"
	"github.com/fuzzyqu/simplified-des/sdes"
)

func main() {

	texto := [8]uint8{0, 0, 1, 0, 1, 0, 0, 0}
	chave := [10]uint8{1, 1, 0, 0, 0, 1, 1, 1, 1, 0}

	cifrado := sdes.Criptografar(texto, chave)
	decifrado := sdes.Descriptografar(cifrado, chave)

	fmt.Println("Original:\t", texto)
	fmt.Println("Cifrado:\t", cifrado)
	fmt.Println("Decifrado:\t", decifrado)

}
