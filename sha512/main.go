package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	data := []byte("Hooray it is 512.")

	// New
	h512 := sha512.New()
	h512.Write(data)
	fmt.Printf("%x\n", h512.Sum(nil))

	// Sum512
	fmt.Printf("%x\n", sha512.Sum512(data))

	// Sum512_256
	h := sha512.New512_256()
	h.Write(data)
	fmt.Printf("%x\n", sha512.New512_256())
}
