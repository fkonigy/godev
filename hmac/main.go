package main

import (
	"crypto/hmac"
	"crypto/md5"
	"fmt"
)

func main() {
	mac1 := getMsg1()
	fmt.Printf("%X\n", mac1)

	mac2 := getMsg2()
	fmt.Printf("%X\n", mac2)

	if hmac.Equal(mac1, mac2) {
		fmt.Println("Messages are identical.")
	} else {
		fmt.Println("Message tampered.")
	}
}

func getMsg1() []byte {
	input := []byte("Secret message.")
	key := []byte("secret key is here.")

	mac := hmac.New(md5.New, key)
	mac.Write(input)

	return mac.Sum(nil)
}

func getMsg2() []byte {
	input := []byte("Secret message.")
	key := []byte("secret key is here.")

	mac := hmac.New(md5.New, key)
	mac.Write(input)

	return mac.Sum(nil)
}
