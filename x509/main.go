package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {

	priv := generateKeys()
	pub := &priv.PublicKey

	fmt.Println(priv)
	fmt.Println()
	fmt.Println(pub)

	cert := createCertificate(pub, priv)

	fmt.Println(cert)

	saveCert(cert)

}

func generateKeys() (priv *rsa.PrivateKey) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return
}

func createCertificate(pub *rsa.PublicKey, priv *rsa.PrivateKey) []byte {

	sn := new(big.Int)
	_, err := fmt.Sscan("1000", sn)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	subject := pkix.Name{
		Country:            []string{"IR"},
		Organization:       []string{"CBI"},
		OrganizationalUnit: []string{"Payment Systems"},
		SerialNumber:       "9000000",
		CommonName:         "CBI Root CA",
	}

	notBefore := time.Now()
	notAfter := time.Time.Add(time.Now(), time.Hour*8760*5)

	template := new(x509.Certificate)
	template.IsCA = true
	template.PublicKey = &pub
	template.SerialNumber = sn
	template.Subject = subject
	template.NotBefore = notBefore
	template.NotAfter = notAfter

	cert, err := x509.CreateCertificate(rand.Reader, template, template, pub, priv)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cert

}

func saveCert(cert []byte) (*os.File, error) {

	file, err := os.OpenFile("root.der", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	file.Write(cert)

	return file, err
}
