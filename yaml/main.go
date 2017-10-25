package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

type MyConfig struct {
	K1 string   `yaml:"K1"`
	K2 int      `yaml:"K2"`
	K3 float32  `yaml:"K3"`
	K4 bool     `yaml:"K4"`
	K5 []string `yaml:"K5"`
	K6 []User   `yaml:"K6"`
}

type User struct {
	Name   string `yaml:"Name"`
	Family string `yaml:"Family"`
	Age    int    `yaml:"Age"`
}

func main() {

	log.SetFlags(log.Lshortfile)

	// Read config file
	data, err := ioutil.ReadFile("yaml_data.yaml")
	checkErr(err, "Erorr opening config file.")

	// print values
	mc := MyConfig{}
	err = yaml.Unmarshal(data, &mc)
	checkErr(err, "Error reading config file.")
	if mc.K4 {
		fmt.Println(mc.K1)
		fmt.Println(mc.K2)
		fmt.Println(mc.K3)

		for _, v := range mc.K5 {
			fmt.Println(v)
		}

		for _, k := range mc.K6 {
			fmt.Println(k.Name, k.Family, k.Age)
		}

	}

}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
