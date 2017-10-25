package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	log.SetFlags(log.Lshortfile)

	fmt.Println("Caller Process ID: ", os.Getpid())
	fmt.Println("Caller Group ID: ", os.Getegid())
	fmt.Println("Caller's parent ID: ", os.Getppid())
	fmt.Println("Caller User ID: ", os.Getuid())

	dir, _ := os.Getwd()
	fmt.Println("Get working directory: ", dir)

	fmt.Println("GOPATH: ", os.Getenv("GOPATH"))
	fmt.Println("GOROOT: ", os.Getenv("GOROOT"))
	fmt.Println("PageSize: ", os.Getpagesize())

	host, _ := os.Hostname()
	fmt.Println("Hostname: ", host)

	fmt.Println("Temp Dir: ", os.TempDir())

	os.Create("create.txt")
	os.Remove("create.txt")

	f1, err := os.Create("File1.txt")
	checkErr(err, "Error creating File1.txt")
	defer f1.Close()

	f2, err := os.Create("File2.txt")
	checkErr(err, "Error creating File2.txt")
	defer f2.Close()

	c1, err := f1.WriteString("Hello Pipe!")
	checkErr(err, "Error writting to File1.txt")
	fmt.Println("Written Count - f1: ", c1)

	f1.Seek(0, 0)

	c2, err := io.Copy(f2, f1)
	checkErr(err, "Error copying File1.txt to File2.txt.")
	fmt.Println("Written Count - f2: ", c2)

	r, w, err := os.Pipe()
	checkErr(err, "Pipe Error!")
	w.WriteString("Pipe R")
	r.WriteString("Pipe W")
	fmt.Println("W Name: ", w.Name())
	fmt.Println("R Name: ", r.Name())

}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
