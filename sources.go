package main

import (
	"bytes"
	"strings"
)

type CardInfo struct {
	ID   string
	Set  string
	SetID string
	Name string
	Quality uint
}

type CustomURL func(info CardInfo) string
type AssetSource struct {
	Quality    uint
	ComposeURL CustomURL
}

var assetList []AssetSource = []AssetSource{
	{
		Quality: 125400,
		ComposeURL: func(info CardInfo) string {
			startIdx := len(info.ID) - 5
			buffer := bytes.NewBufferString("http://netrunnerdb.com/web/bundles/netrunnerdbcards/images/cards/en/")
			buffer.WriteString(info.ID[startIdx:])
			buffer.WriteString(".png")
			return buffer.String()
		},
	},
	{
		Quality: 58870,
		ComposeURL: func(info CardInfo) string {
			buffer := bytes.NewBufferString("http://www.cardgamedb.com/forums/uploads/an/")
			buffer.WriteString("ffg_")
			buffer.WriteString(strings.Replace(strings.ToLower(info.Name), " ", "-", -1))
			buffer.WriteString("-")
			buffer.WriteString(strings.Replace(strings.ToLower(info.Set), " ", "-", -1))
			buffer.WriteString(".png")
			return buffer.String()
		},
	},
}