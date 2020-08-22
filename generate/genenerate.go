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
		TName string
	}

	type EnumType struct {
		Name    string
		Members []EnumMember
	}

	for _, et := range schema.DataServices.Schema.EnumType {
		p := filepath.Join("..", "graph", fmt.Sprintf("%sEnum.go", strings.Title(et.Name)))
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
				Name:  m.Name,
				TName: strings.Title(m.Name),
				Value: m.Value,
			})
		}
		enumTemplate.ExecuteTemplate(out, "enum.go.templ", e)
		out.Close()
	}

	type Property struct {
		Name           string
		Type           string
		IsCollection   bool
		CollectionType string
	}

	type EntityType struct {
		Name       string
		BaseType   string
		OpenType   string
		Properties []Property
	}

	for _, etp := range schema.DataServices.Schema.EntityType {
		p := filepath.Join("..", "graph", fmt.Sprintf("%sTypes.go", strings.Title(etp.Name)))
		out, err := os.Create(p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		t := EntityType{
			Name:       etp.Name,
			BaseType:   etp.BaseType,
			OpenType:   etp.OpenType,
			Properties: make([]Property, 0),
		}
		for _, p := range etp.Property {
			property := Property{
				Name: p.Name,
			}
			property.Type = XmlToGoType(p.Type)

			t.Properties = append(t.Properties, property)
		}
		enumTemplate.ExecuteTemplate(out, "entity.type.go.templ", t)
		out.Close()
	}
}

func XmlToGoType(t string) string {
	if strings.Contains(t, "Collection") {
		s := strings.Trim(t, "Collection(")
		s1 := strings.Trim(s, ")")

		split := strings.Split(s1, ".")
		return fmt.Sprintf("[]%s", strings.Title(split[len(split)-1]))

	} else if strings.Contains(t, "microsoft.graph.") {
		split := strings.Split(t, ".")
		return fmt.Sprintf("*%s", strings.Title(split[len(split)-1]))
	} else if strings.Contains(t, "Edm.") {
		split := strings.Split(t, ".")
		return fmt.Sprintf("*%s", strings.Title(split[len(split)-1]))
	} else {
		panic(fmt.Sprintf("Handle tye type %s", t))
	}
}
