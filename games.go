package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
)

type Game struct {
	Name        string
	ID          string
	IgnoreSets  []string
	IgnoreCards []string
	ComposeURL  func(info CardInfo) string
}

var gameList = []Game{
	netrunner,
	thrones,
}

var netrunner Game = Game{
	Name:        "Netrunner",
	ID:          "0f38e453-26df-4c04-9d67-6d43de939c77",
	IgnoreSets:  []string{"21bf9e05-fb23-4b1d-b89a-398f671f5999"},
	IgnoreCards: []string{"bc0f047c-01b1-427f-a439-d451eda00000"},
	ComposeURL: func(info CardInfo) string {
		startIdx := len(info.ID) - 5
		buffer := bytes.NewBufferString("http://netrunnerdb.com/bundles/netrunnerdbcards/images/cards/en/")
		buffer.WriteString(info.ID[startIdx:])
		buffer.WriteString(".png")
		return buffer.String()
	},
}

var thrones Game = Game{
	Name:        "AGoT 2.0",
	ID:          "30c200c9-6c98-49a4-a293-106c06295c05",
	IgnoreSets:  []string{"d2e1abfc-fead-4067-a3b7-14973a19ca21"},
	IgnoreCards: []string{},
	ComposeURL: func(info CardInfo) string {
		return fmt.Sprintf("http://www.thronesdb.com/bundles/cards/%s%03s.png", info.Set[:2], info.Number)
	},
}

//Get important OCTGN Netrunner directory paths
func getPaths(gInfo Game) (setPath string, imgPath string) {
	curUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	installPath := path.Join(curUser.HomeDir, "Documents", "OCTGN")
	if *forceCWD {
		installPath, _ = os.Getwd()
	}

	setPath = path.Join(installPath, "GameDatabase", gInfo.ID, "Sets")
	imgPath = path.Join(installPath, "ImageDatabase", gInfo.ID, "Sets")
	return
}
