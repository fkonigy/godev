package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	input := []byte("This time weuse SHA256.")

	// New
	h := sha256.New()
	h.Write(input)
	fmt.Printf("%x\n", h.Sum(nil))

	// Sum256
	fmt.Printf("%x\n", sha256.Sum256(input))

	// New224
	h224 := sha256.New224()
	h224.Write(input)
	fmt.Printf("%x\n", h224.Sum(nil))

	// Sum224
	fmt.Printf("%x\n", sha256.Sum224(input))

}
