package main

//go:generate go run .

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello world!!")

	xmlFile, err := os.Open("generate/v1.0.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var schema Edmx
	err = xml.Unmarshal(byteValue, &schema)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, et := range schema.DataServices.Schema.EnumType {
		fmt.Printf("-----> %v\n", et.Name)
		for _, m := range et.Member {
			fmt.Printf("----------> %v", m.Name)
		}
	}
}
