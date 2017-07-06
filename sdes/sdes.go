package sdes

import (
	"fmt"
	"log"
	"strconv"
)

// todo: Unir as funções de criptografar e descriptografar ?
// todo: Eu gosto da clareza de deixa-las separadas, mas repeat yourself é sempre sad.

// Criptografar criptografa um texto utilizando a cifra des simplificada.
func Criptografar(texto [8]uint8, chave [10]uint8) [8]uint8 {

	// Sanity check xD
	// Isso é meio idiota, arrays vão estar com o tamanho certo sempre.
	if len(texto) != 8 {
		log.Fatal("O texto precisa ter 8 bits :c")
	}
	if len(chave) != 10 {
		log.Fatal("A chave precisa ter 10 bits :c")
	}

	// Primeiro vamos criar as subchaves.
	chave = p10(chave)
	chave = shift(chave, 1)
	k1 := p8(chave)
	chave = shift(chave, 2)
	k2 := p8(chave)

	// Agora Criptografar o texto!
	textoCifrado := ip(texto)
	textoCifrado = fk(textoCifrado, k1)
	textoCifrado = sw(textoCifrado)
	textoCifrado = fk(textoCifrado, k2)
	textoCifrado = ipi(textoCifrado)

	return textoCifrado
}

// Descriptografar descriptografa um texto cifrado pela cifra des simplificada.
func Descriptografar(texto [8]uint8, chave [10]uint8) [8]uint8 {

	// Primeiro vamos criar as subchaves.
	chave = p10(chave)
	chave = shift(chave, 1)
	k1 := p8(chave)
	chave = shift(chave, 2)
	k2 := p8(chave)

	// Agora descriptografar o texto!
	// A única diferença é a ordem de k1/k2.
	textoDecifrado := ip(texto)
	textoDecifrado = fk(textoDecifrado, k2)
	textoDecifrado = sw(textoDecifrado)
	textoDecifrado = fk(textoDecifrado, k1)
	textoDecifrado = ipi(textoDecifrado)

	return textoDecifrado
}

// Funções relacionadas ao texto.
func ip(texto [8]uint8) [8]uint8 {
	var resultado [8]uint8
	resultado[0], resultado[1], resultado[2], resultado[3] = texto[1], texto[5], texto[2], texto[0]
	resultado[4], resultado[5], resultado[6], resultado[7] = texto[3], texto[7], texto[4], texto[6]

	fmt.Println("ip: ", resultado)
	return resultado
}

func fk(texto, sk [8]uint8) [8]uint8 {

	var resultado [8]uint8
	var lTexto [4]uint8
	var rTexto [4]uint8

	// Separamos em um par de 4 bits.
	copy(lTexto[:], texto[:4])
	copy(rTexto[:], texto[4:])

	// Primeiro aplicamos f nos bits a direita.
	fBits := f(rTexto, sk)

	// Agora fazemos um XORzinho do resultado de f com os bits da esquerda.
	for i := range lTexto {
		lTexto[i] ^= fBits[i]
	}

	// Pronto, agora colocamos os bits no lugar.
	// Juntando as duas metades.
	copy(resultado[:4], lTexto[:])
	copy(resultado[4:], rTexto[:])

	fmt.Println("fk: ", resultado)
	return resultado
}

func sw(texto [8]uint8) [8]uint8 {
	var resultado [8]uint8

	// Troca as metades de lugar.
	copy(resultado[4:], texto[:4])
	copy(resultado[:4], texto[4:])

	fmt.Println("sw: ", resultado)
	return resultado
}

func ipi(texto [8]uint8) [8]uint8 {
	var resultado [8]uint8
	resultado[0], resultado[1], resultado[2], resultado[3] = texto[3], texto[0], texto[2], texto[4]
	resultado[4], resultado[5], resultado[6], resultado[7] = texto[6], texto[1], texto[7], texto[5]

	fmt.Println("ip^-1: ", resultado)
	return resultado
}

// Funções relacionadas a chave.
func p10(chave [10]uint8) [10]uint8 {
	var resultado [10]uint8
	copy(resultado[:], chave[:])
	resultado[0], resultado[1], resultado[2], resultado[3], resultado[4] = chave[2], chave[4], chave[1], chave[6], chave[3]
	resultado[5], resultado[6], resultado[7], resultado[8], resultado[9] = chave[9], chave[0], chave[8], chave[7], chave[5]

	//fmt.Println("p10: ", resultado)
	return resultado
}

func shift(chave [10]uint8, n int) [10]uint8 {

	var bitsEsq [5]uint8
	var bitsDir [5]uint8

	// Separamos em dois a chave.
	copy(bitsEsq[:], chave[:5])
	copy(bitsDir[:], chave[5:])

	// Certo, aqui é feita uma rotação a esquerda (n vezes)
	// Cada metade é rotacionada separadamente.
	for i := 0; i < n; i++ {
		e := bitsEsq[0]
		d := bitsDir[0]
		copy(bitsEsq[:], bitsEsq[1:])
		copy(bitsDir[:], bitsDir[1:])
		bitsEsq[len(bitsEsq)-1] = e
		bitsDir[len(bitsDir)-1] = d
	}

	// Juntando as partes rotacionadas.
	copy(chave[:5], bitsEsq[:])
	copy(chave[5:], bitsDir[:])

	//fmt.Println("shift: ", chave)
	return chave
}

func p8(chave [10]uint8) [8]uint8 {
	var resultado [8]uint8
	resultado[0], resultado[1], resultado[2], resultado[3] = chave[5], chave[2], chave[6], chave[3]
	resultado[4], resultado[5], resultado[6], resultado[7] = chave[7], chave[4], chave[9], chave[8]

	//fmt.Println("p8: ", resultado)
	return resultado
}

// Funções Extras
func f(n [4]uint8, sk [8]uint8) [4]uint8 {

	var resultado [4]uint8

	// Executa a E/P, expansão/permutação.
	var ep [8]uint8
	ep[0], ep[1], ep[2], ep[3] = n[3], n[0], n[1], n[2]
	ep[4], ep[5], ep[6], ep[7] = n[1], n[2], n[3], n[0]

	// XOR com a subchave
	for i := range ep {
		ep[i] ^= sk[i]
	}

	s0 := [4][4]uint8{{1, 0, 3, 2}, {3, 2, 1, 0}, {0, 2, 1, 3}, {3, 1, 3, 2}}
	s1 := [4][4]uint8{{0, 1, 2, 3}, {2, 0, 1, 3}, {3, 0, 1, 0}, {2, 1, 0, 3}}

	// Concatena os bits da linha 0 e transforma em inteiros.
	row0, _ := strconv.ParseInt(fmt.Sprintf("%b%b", ep[0], ep[3]), 2, 32)
	col0, _ := strconv.ParseInt(fmt.Sprintf("%b%b", ep[1], ep[2]), 2, 32)

	// Concatena os bits da linha 1 e transforma em inteiros.
	row1, _ := strconv.ParseInt(fmt.Sprintf("%b%b", ep[4], ep[7]), 2, 32)
	col1, _ := strconv.ParseInt(fmt.Sprintf("%b%b", ep[5], ep[6]), 2, 32)

	// Pega a representação em bits dos números escolhidos das s-boxs (concatenados).
	resultadoBits := fmt.Sprintf("%02b%02b", s0[row0][col0], s1[row1][col1])

	// Strings são más, voltamos para uint8.
	for i := range resultadoBits {
		resultado[i] = uint8(resultadoBits[i] - '0')
	}

	// Executa P4 :)
	resultado[0], resultado[1], resultado[3] = resultado[1], resultado[3], resultado[0]

	//fmt.Println("f: ", resultado)
	return resultado
}
