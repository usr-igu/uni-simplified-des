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
	p10 = lr(p10)
	p8 := p8(p10)
	return p8
}

// p10 permuta os bits da chave (10) como especificado no documento S-DES.
func p10(key []byte) []byte {
	p10 := append([]byte(nil), key...)
	p10[0], p10[1], p10[2], p10[3], p10[4] = p10[2], p10[4], p10[1], p10[7], p10[3]
	p10[5], p10[6], p10[7], p10[8], p10[9] = p10[9], p10[0], p10[8], p10[7], p10[5]
	return p10
}

// lr separa a chave (10) em 2 grupos de 5 bits e respectivamente executa uma rotação a esquerda.
func lr(key []byte) []byte {
	fh := append([]byte(nil), key[:len(key)/2]...) // Cria uma cópia da primeira metade dos bits.
	sh := append([]byte(nil), key[len(key)/2:]...) // Cria uma cópia da segunda metade dos bits.
	fh = append(fh[1:], fh[0])                     // Circular left shift na primeira metade.
	sh = append(sh[1:], sh[0])                     // Circular left shift na segunda metade.
	lrk := append(fh, sh...)
	return lrk
}

// p8 permuta os bits da chave (10) em um novo conjunto menor de bits (8).
func p8(key []byte) []byte {
	p8 := make([]byte, 8)
	p8[0], p8[1], p8[2], p8[3] = key[5], key[2], key[6], key[3]
	p8[4], p8[5], p8[6], p8[7] = key[7], key[4], key[9], key[8]
	return p8
}
