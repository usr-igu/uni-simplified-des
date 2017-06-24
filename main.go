package main

import (
	"fmt"
	"github.com/fuzzyquanta/sdes/sdes"
)

func main() {
	fmt.Printf("%+v\n", sdes.SimplifiedDES([]byte{1, 0, 1, 0, 1, 0, 1, 0}, []byte{1, 0, 1, 0, 0, 0, 0, 0, 1, 0}))
}
