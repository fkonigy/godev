package main

import "fmt"

func main() {
	name := "Hi"
	n0, err := fmt.Printf("He says %s", name)
	// go does not complain about err being used again, because there is a new
	// varibale declaration n1 and err is not redeclared again.
	n1, err := fmt.Printf("She says %s", name)

	fmt.Printf("We have %d and %d. %s\n", n0, n1, err)

	// but this time go complains about err, because there is no new variable
	// declaration.
	_, err := fmt.Printf("He says %s", name)
	fmt.Printf("%s\n", err)

}
