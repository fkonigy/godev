package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	input := "Hello Crypto."
	h := md5.New()
	io.WriteString(h, input)
	fmt.Printf("%x\n", h.Sum(nil))

	io.WriteString(os.Stdout, "Hello StdOut!\n")
	fmt.Println("Hello Println!\n")

	// This time we use Sum method directly.
	data := []byte("Hello Crypto.")
	fmt.Printf("%x\n", md5.Sum(data))
}
