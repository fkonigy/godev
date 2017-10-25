package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	input := "Hash me with SHA1 please."

	h := sha1.New()
	io.WriteString(h, input)

	fmt.Printf("%x\n", h.Sum(nil))

	data := []byte("Hash me with SHA1 please.")
	fmt.Printf("%x\n", sha1.Sum(data))

}
