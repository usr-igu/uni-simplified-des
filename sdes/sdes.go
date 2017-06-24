package sdes

import (
	"log"
)

func SimplifiedDES(plaintext []byte, key []byte) []byte {

	if len(plaintext) != 8 {
		log.Fatalf("Texto a ser criptografado precisa ter 8 bits: len(plaintext) == %d.", len(plaintext))
	}

	if len(key) != 10 {
		log.Fatalf("Chave de criptografia precisa ter 10 bits: len(plaintext) == %d.", len(plaintext))
	}

	p10 := p10(key)
	p10 = circularLeftShift(p10)
	p8 := p8(p10)

	return p8
}

func p10(p10 []byte) []byte {
	p10[0], p10[1], p10[2], p10[3], p10[4] = p10[2], p10[4], p10[1], p10[7], p10[3]
	p10[5], p10[6], p10[7], p10[8], p10[9] = p10[9], p10[0], p10[8], p10[7], p10[5]
	return p10
}

func circularLeftShift(p10 []byte) []byte {

	fh := append([]byte(nil), p10[:len(p10)/2]...) // Primeira metade dos bits.
	sh := append([]byte(nil), p10[len(p10)/2:]...) // Segunda metade dos bits.

	var x, y byte

	x, fh = fh[0], fh[1:]
	fh = append(fh, x)

	y, sh = sh[0], sh[1:]
	sh = append(sh, y)

	p10 = append(fh, sh...)

	return p10
}

func p8(p10 []byte) []byte {
	p8 := make([]byte, 8)
	p8[0], p8[1], p8[2], p8[3] = p10[5], p10[2], p10[6], p10[3]
	p8[4], p8[5], p8[6], p8[7] = p10[7], p10[4], p10[9], p10[8]
	return p8
}
