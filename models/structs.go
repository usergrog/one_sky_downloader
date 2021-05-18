package models

import "encoding/xml"

type Translation struct {
	Translation interface{}
}

type XMLString struct {
	XMLName   xml.Name `xml:"string"̀`
	Name      string   `xml:"name,attr"`
	Formatted bool     `xml:"formatted,attr"̀`
	Value     string   `xml:",chardata"`
}

type XMLResources struct {
	XMLName    xml.Name    `xml:"resources"̀`
	XMLStrings []XMLString `xml:"string"`
}

type Config struct {
	ApiKey    string
	ApiSecret string
	ProjectId string
	FileName  string
	Format    string
}
