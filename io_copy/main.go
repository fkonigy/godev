package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if c, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Count: ", c)
	}

}
