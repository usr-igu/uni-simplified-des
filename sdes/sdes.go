package sdes

func SDES(texto [8]uint8, chave [10]uint8) [8]uint8 {

	// Primeiro vamos criar as subchaves.
	chave = p10(chave)
	chave = shift(chave)
	k1 := p8(chave)
	chave = shift(chave)
	k2 := p8(chave)

	// Agora Criptografar o texto!
	textoCifrado := ip(texto)
	textoCifrado = fk(textoCifrado, k1)
	textoCifrado = sw(textoCifrado)
	textoCifrado = fk(textoCifrado, k2)
	textoCifrado = ipi(textoCifrado)

	return textoCifrado
}

// Funções relacionadas ao plaintext.
func ip(texto [8]uint8) [8]uint8 {
	var aux [8]uint8
	aux[0], aux[1], aux[2], aux[3] = texto[1], texto[5], texto[2], texto[0]
	aux[4], aux[5], aux[6], aux[7] = texto[3], texto[7], texto[4], texto[6]
	return aux
}

// todo
func fk() {

}

func sw(texto [8]uint8) [8]uint8 {
	var aux [8]uint8
	copy(aux[4:], texto[:4])
	copy(aux[:4], texto[4:])
	return aux
}

func ipi(texto [8]uint8) [8]uint8 {
	var aux [8]uint8
	aux[0], aux[1], aux[2], aux[3] = texto[3], texto[0], texto[2], texto[4]
	aux[4], aux[5], aux[6], aux[7] = texto[6], texto[1], texto[7], texto[5]
	return aux
}

// Funções relacionadas a chave.
func p10(chave [10]uint8) [10]uint8 {
	var aux [10]uint8
	copy(aux[:], chave[:])
	aux[0], aux[1], aux[2], aux[3], aux[4] = chave[2], chave[4], chave[1], chave[6], chave[3]
	aux[5], aux[6], aux[7], aux[8], aux[9] = chave[9], chave[0], chave[8], chave[7], chave[5]
	return aux
}

// todo
func shift() {

}

func p8(chave [10]uint8) [8]uint8 {
	var aux [8]uint8
	aux[0], aux[1], aux[2], aux[3] = chave[5], chave[2], chave[6], chave[3]
	aux[4], aux[5], aux[6], aux[7] = chave[7], chave[4], chave[9], chave[8]
	return aux
}

// Funções Extras
