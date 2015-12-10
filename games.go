package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"strconv"
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
	wh40kconquest,
}

var netrunner = Game{
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

var thrones = Game{
	Name:        "AGoT 2.0",
	ID:          "30c200c9-6c98-49a4-a293-106c06295c05",
	IgnoreSets:  []string{"d2e1abfc-fead-4067-a3b7-14973a19ca21"},
	IgnoreCards: []string{},
	ComposeURL: func(info CardInfo) string {
		return fmt.Sprintf("http://www.thronesdb.com/bundles/cards/%s%03s.png", info.Set[:2], info.Number)
	},
}

// Will not download Core Set (higher res images available at http://octgngames.com/wh40kc/ already)
var wh40kconquest = Game{
	Name:        "Warhammer40k Conquest",
	ID:          "af04f855-58c4-4db3-a191-45fe33381679",
	IgnoreSets:  []string{"cdba7854-4c22-48f3-b388-74ca361b05d9", "35c6df08-5a89-47bb-b8f3-624bcd8d9d43"},
	IgnoreCards: []string{},
	ComposeURL: func(info CardInfo) string {
		return fmt.Sprintf("http://s3.amazonaws.com/LCG/40kconquest/med_WHK%s_%s.jpg", wh40kSubset(info), info.Number)
	},
}

// The Warlord Cycle is broken into subsets in the image database, while the GameDatabase definition just has them
// all defined as a single set. Break them down by card number.
func wh40kSubset(info CardInfo) (subset string) {
	cardNum, _ := strconv.Atoi(info.Number)

	switch {
	default:
		return "01"
	case info.Set == "01-Core Set":
		return "01"
	case info.Set == "02- Warlord Cycle":
		if cardNum < 23 {
			return "02"
		} else if cardNum < 45 {
			return "03"
		} else if cardNum < 67 {
			return "04"
		} else if cardNum < 89 {
			return "05"
		} else if cardNum < 111 {
			return "06"
		}
		return "07"
	case info.Set == "03-The Great Devourer":
		return "08"
	}

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
