package main

//go:generate go run .

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Hello world!!")

	xmlFile, err := os.Open("v1.0.xml")
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

	enumTemplate, err := template.ParseFiles("templates/enum.go.templ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	type EnumMember struct {
		Name  string
		Value string
	}

	type EnumType struct {
		Name    string
		Members []EnumMember
	}

	for _, et := range schema.DataServices.Schema.EnumType {
		p := filepath.Join("..", "graph", fmt.Sprintf("%s.go", et.Name))
		out, err := os.Create(p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		e := EnumType{
			Name:    strings.Title(et.Name),
			Members: make([]EnumMember, 0),
		}
		for _, m := range et.Member {
			e.Members = append(e.Members, EnumMember{
				Name:  strings.Title(m.Name),
				Value: m.Value,
			})
		}
		enumTemplate.ExecuteTemplate(out, "enum.go.templ", e)
		out.Close()
	}
}
