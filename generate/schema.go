package main

import "encoding/xml"

type Edmx struct {
	XMLName      xml.Name `xml:"Edmx"`
	Text         string   `xml:",chardata"`
	Version      string   `xml:"Version,attr"`
	Edmx         string   `xml:"edmx,attr"`
	DataServices struct {
		Text   string `xml:",chardata"`
		Schema struct {
			Text      string `xml:",chardata"`
			Namespace string `xml:"Namespace,attr"`
			Xmlns     string `xml:"xmlns,attr"`
			EnumType  []struct {
				Text    string `xml:",chardata"`
				Name    string `xml:"Name,attr"`
				IsFlags string `xml:"IsFlags,attr"`
				Member  []struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"Name,attr"`
					Value string `xml:"Value,attr"`
				} `xml:"Member"`
			} `xml:"EnumType"`
			EntityType []struct {
				Text      string `xml:",chardata"`
				Name      string `xml:"Name,attr"`
				Abstract  string `xml:"Abstract,attr"`
				BaseType  string `xml:"BaseType,attr"`
				OpenType  string `xml:"OpenType,attr"`
				HasStream string `xml:"HasStream,attr"`
				Key       struct {
					Text        string `xml:",chardata"`
					PropertyRef struct {
						Text string `xml:",chardata"`
						Name string `xml:"Name,attr"`
					} `xml:"PropertyRef"`
				} `xml:"Key"`
				Property []struct {
					Text     string `xml:",chardata"`
					Name     string `xml:"Name,attr"`
					Type     string `xml:"Type,attr"`
					Nullable string `xml:"Nullable,attr"`
				} `xml:"Property"`
				NavigationProperty []struct {
					Text           string `xml:",chardata"`
					Name           string `xml:"Name,attr"`
					Type           string `xml:"Type,attr"`
					ContainsTarget string `xml:"ContainsTarget,attr"`
					Nullable       string `xml:"Nullable,attr"`
				} `xml:"NavigationProperty"`
			} `xml:"EntityType"`
			ComplexType []struct {
				Text     string `xml:",chardata"`
				Name     string `xml:"Name,attr"`
				OpenType string `xml:"OpenType,attr"`
				BaseType string `xml:"BaseType,attr"`
				Abstract string `xml:"Abstract,attr"`
				Property []struct {
					Text     string `xml:",chardata"`
					Name     string `xml:"Name,attr"`
					Type     string `xml:"Type,attr"`
					Nullable string `xml:"Nullable,attr"`
				} `xml:"Property"`
			} `xml:"ComplexType"`
			Action []struct {
				Text          string `xml:",chardata"`
				Name          string `xml:"Name,attr"`
				IsBound       string `xml:"IsBound,attr"`
				EntitySetPath string `xml:"EntitySetPath,attr"`
				Parameter     []struct {
					Text     string `xml:",chardata"`
					Name     string `xml:"Name,attr"`
					Type     string `xml:"Type,attr"`
					Nullable string `xml:"Nullable,attr"`
					Unicode  string `xml:"Unicode,attr"`
				} `xml:"Parameter"`
				ReturnType struct {
					Text     string `xml:",chardata"`
					Type     string `xml:"Type,attr"`
					Nullable string `xml:"Nullable,attr"`
					Unicode  string `xml:"Unicode,attr"`
				} `xml:"ReturnType"`
			} `xml:"Action"`
			Function []struct {
				Text          string `xml:",chardata"`
				Name          string `xml:"Name,attr"`
				IsBound       string `xml:"IsBound,attr"`
				IsComposable  string `xml:"IsComposable,attr"`
				EntitySetPath string `xml:"EntitySetPath,attr"`
				Parameter     []struct {
					Text     string `xml:",chardata"`
					Name     string `xml:"Name,attr"`
					Type     string `xml:"Type,attr"`
					Nullable string `xml:"Nullable,attr"`
					Unicode  string `xml:"Unicode,attr"`
				} `xml:"Parameter"`
				ReturnType struct {
					Text     string `xml:",chardata"`
					Type     string `xml:"Type,attr"`
					Nullable string `xml:"Nullable,attr"`
					Unicode  string `xml:"Unicode,attr"`
				} `xml:"ReturnType"`
			} `xml:"Function"`
			Term []struct {
				Text      string `xml:",chardata"`
				Name      string `xml:"Name,attr"`
				Type      string `xml:"Type,attr"`
				AppliesTo string `xml:"AppliesTo,attr"`
			} `xml:"Term"`
			EntityContainer struct {
				Text      string `xml:",chardata"`
				Name      string `xml:"Name,attr"`
				EntitySet []struct {
					Text                      string `xml:",chardata"`
					Name                      string `xml:"Name,attr"`
					EntityType                string `xml:"EntityType,attr"`
					NavigationPropertyBinding []struct {
						Text   string `xml:",chardata"`
						Path   string `xml:"Path,attr"`
						Target string `xml:"Target,attr"`
					} `xml:"NavigationPropertyBinding"`
				} `xml:"EntitySet"`
				Singleton []struct {
					Text                      string `xml:",chardata"`
					Name                      string `xml:"Name,attr"`
					Type                      string `xml:"Type,attr"`
					NavigationPropertyBinding []struct {
						Text   string `xml:",chardata"`
						Path   string `xml:"Path,attr"`
						Target string `xml:"Target,attr"`
					} `xml:"NavigationPropertyBinding"`
				} `xml:"Singleton"`
			} `xml:"EntityContainer"`
			Annotations []struct {
				Text       string `xml:",chardata"`
				Target     string `xml:"Target,attr"`
				Annotation []struct {
					Text   string `xml:",chardata"`
					Term   string `xml:"Term,attr"`
					Bool   string `xml:"Bool,attr"`
					String string `xml:"String,attr"`
					Record struct {
						Text          string `xml:",chardata"`
						PropertyValue struct {
							Text       string `xml:",chardata"`
							Property   string `xml:"Property,attr"`
							Bool       string `xml:"Bool,attr"`
							EnumMember string `xml:"EnumMember"`
						} `xml:"PropertyValue"`
					} `xml:"Record"`
					EnumMember string `xml:"EnumMember"`
				} `xml:"Annotation"`
			} `xml:"Annotations"`
		} `xml:"Schema"`
	} `xml:"DataServices"`
}
