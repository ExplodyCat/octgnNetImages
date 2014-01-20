package main

import (
	"log"
	"encoding/xml"
	"io/ioutil"
)


type NetSet struct{
	XMLName xml.Name `xml:"set"`
	Name string `xml:"name,attr"`
	ID string `xml:"id,attr"`
	Cards NetCards `xml:"cards"`
}

type NetCards struct{	
	Cards []NetCard `xml:"card"`
}
type NetCard struct{
	ID string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

func parseSetXML(xmlPath string) []CardInfo {
	v:= new(NetSet)
	data, err := ioutil.ReadFile(xmlPath)
	if err !=nil{
		log.Fatal(err)
	}
	
	err = xml.Unmarshal(data,&v)
	if err!=nil{
		log.Fatal(err)
	}
	results := []CardInfo{}
	for _, cur:= range v.Cards.Cards{
			newItem := CardInfo{
				ID:cur.ID,
				Name:cur.Name,
				Set:v.Name,
				SetID:v.ID,
				}
			results = append(results, newItem)
	}	
	return results
}
