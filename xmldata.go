package main

import (
	"encoding/xml"
	"io/ioutil"
	"regexp"
)

type CardInfo struct {
	ID     string
	Set    string
	SetID  string
	Name   string
	Number string
}

type set struct {
	XMLName xml.Name `xml:"set"`
	Name    string   `xml:"name,attr"`
	ID      string   `xml:"id,attr"`
	Cards   []card   `xml:"cards>card"`
}

type card struct {
	ID      string `xml:"id,attr"`
	Name    string `xml:"name,attr"`
	Details []prop `xml:"property"`
}
type prop struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

var (
	re = regexp.MustCompile("[0-9]+")
)

//Parse a set xml file and return slice of cards
func parseSetXML(xmlPath string) (results []CardInfo, err error) {
	//read data from xml file
	data, err := ioutil.ReadFile(xmlPath)
	if err != nil {
		return results, err
	}

	//parse that xml data
	v := new(set)
	err = xml.Unmarshal(data, &v)
	if err != nil {
		return results, err
	}

	//parse out all the cards
	for _, cur := range v.Cards {
		cNumber := ""
		for _, det := range cur.Details {
			if det.Name == "CardNumber" {
				cNumber = re.FindString(det.Value)
			}
		}
		newItem := CardInfo{
			ID:     cur.ID,
			Name:   cur.Name,
			Set:    v.Name,
			SetID:  v.ID,
			Number: cNumber,
		}
		results = append(results, newItem)
	}
	return
}
