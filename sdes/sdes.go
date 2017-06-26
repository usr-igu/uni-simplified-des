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

// p10 permuta os bits da chave (10).
//
// A permutação é especificada no documento S-DES.
func p10(key []byte) []byte {
	p10 := make([]byte, len(key))
	copy(p10, key) // Copia os bits da chave.
	p10[0], p10[1], p10[2], p10[3], p10[4] = p10[2], p10[4], p10[1], p10[7], p10[3]
	p10[5], p10[6], p10[7], p10[8], p10[9] = p10[9], p10[0], p10[8], p10[7], p10[5]
	return p10
}

// p8 permuta os bits da chave (10) em um novo conjunto menor de bits (8).
//
// A permutação é especificada no documento S-DES.
func p8(key []byte) []byte {
	p8 := make([]byte, 8)
	p8[0], p8[1], p8[2], p8[3] = key[5], key[2], key[6], key[3]
	p8[4], p8[5], p8[6], p8[7] = key[7], key[4], key[9], key[8]
	return p8
}

// ip permuta os bits de plaintext (8).
//
// A permutação é especificada no documento S-DES.
func ip(txt []byte) []byte {
	nTxt := make([]byte, 8)
	nTxt[0], nTxt[1], nTxt[2], nTxt[3] = txt[1], txt[5], txt[2], txt[0]
	nTxt[4], nTxt[5], nTxt[6], nTxt[7] = txt[3], txt[7], txt[4], txt[6]
	return nTxt
}

// ipi permuta os bits de plaintext (8) de forma inversa a ip.
//
// A permutação é especificada no documento S-DES.
func ipi(ptxt []byte) []byte {
	ptxtn := make([]byte, 8)
	ptxtn[0], ptxtn[1], ptxtn[2], ptxtn[3] = ptxt[3], ptxt[0], ptxt[2], ptxt[4]
	ptxtn[4], ptxtn[5], ptxtn[6], ptxtn[7] = ptxt[6], ptxt[1], ptxt[7], ptxt[5]
	return ptxtn
}

// lr separa a chave (10) em 2 grupos de 5 bits e respectivamente executa uma rotação a esquerda.
func lr(key []byte) []byte {
	fh := make([]byte, len(key)/2)
	sh := make([]byte, len(key)/2)
	copy(fh, key[:len(key)/2]) // Copia a primeira metade dos bits.
	copy(sh, key[len(key)/2:]) // Copia a cópia da segunda metade dos bits.
	fh = append(fh[1:], fh[0]) // Circular left shift na primeira metade.
	sh = append(sh[1:], sh[0]) // Circular left shift na segunda metade.
	return append(fh, sh...)
}

//// todo: Terminar fk.
//// fk
//func fk(key []byte) []byte {
//	fh := append([]byte(nil), key[:len(key)/2]...) // Cria uma cópia da primeira metade dos bits.
//	sh := append([]byte(nil), key[len(key)/2:]...) // Cria uma cópia da segunda metade dos bits.
//}

// sw troca a primeira metade dos bits de key pela segunda.
func sw(key []byte) []byte {
	fh := make([]byte, len(key)/2)
	sh := make([]byte, len(key)/2)
	copy(fh, key[:len(key)/2]) // Copia a primeira metade dos bits.
	copy(sh, key[len(key)/2:]) // Copia a cópia da segunda metade dos bits.
	return append(sh, fh...)
}
