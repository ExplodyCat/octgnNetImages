package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
)

type Task struct {
	Dst  string
	Src  string
	Card CardInfo
}

const (
	consumeThreads int = 4
	chanSize       int = 60
)

var (
	wGroup        sync.WaitGroup
	wChan         = make(chan Task, chanSize)
	forceDownload = flag.Bool("f", false, "Force download of all images")
	forceCWD      = flag.Bool("c", false, "Treat current working dir as root OCTGN. Good for nonstandard OCTGN locations.")
)

//TODO: provide a flag to specify download for only a specific game

func main() {
	flag.Parse()

	for i := 0; i < consumeThreads; i++ {
		wGroup.Add(1)
		go consumer()
	}
	for _, gInfo := range gameList {
		producer(gInfo)
	}
	close(wChan)
	wGroup.Wait()
}

func searchList(target string, list []string) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

//Loads wChan with list of cards to consider downloading
func producer(gInfo Game) {
	setPath, imgPath := getPaths(gInfo)
	//If the game isn't installed, return
	if _, err := os.Stat(setPath); os.IsNotExist(err) {
		fmt.Printf("Game not found: %s\n\t%s\n\n", gInfo.Name, setPath)
		return
	}
	//get list of set directories
	setDirs, err := ioutil.ReadDir(setPath)
	if err != nil {
		fmt.Printf("Game Error: %s - %s\n", gInfo.Name, err)
		return
	}

	for _, curSet := range setDirs {

		//Skip over ignored sets
		if searchList(curSet.Name(), gInfo.IgnoreSets) {
			continue
		}

		//Get the collection of cards from the set
		setFile := fmt.Sprintf("%s/%s/set.xml", setPath, curSet.Name())
		setColl, err := parseSetXML(setFile)
		if err != nil {
			fmt.Printf("Error loading set file: %s\n\t%s\n", setFile, err)
			continue
		}

		for _, curCard := range setColl {
			//Skip promos/undownloadable cards
			if searchList(curCard.ID, gInfo.IgnoreCards) {
				continue
			}

			dst := path.Join(imgPath, curCard.SetID, "Cards", curCard.ID+".png")
			src := gInfo.ComposeURL(curCard)

			//Only download files that don't exist
			if _, err := os.Stat(dst); os.IsNotExist(err) || *forceDownload {
				wChan <- Task{dst, src, curCard}
			}
		}
	}
}

//Waits for and processes download tasks
func consumer() {
	defer wGroup.Done()
	for curTask := range wChan {
		fmt.Printf("Attempting download: %s - %s\n", curTask.Card.Set, curTask.Card.Name)
		if err := doDownload(curTask.Src, curTask.Dst); err != nil {
			fmt.Printf("Failed Get: %s - %s\n\t%s\n", curTask.Card.Set, curTask.Card.Name, err)
		}
	}
}

//Download contents of url to target path
func doDownload(src string, dst string) (err error) {
	//Get the url contents
	resp, err := http.Get(src)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to get url: HTTP:%d:%s\n\t%s\n", resp.StatusCode, http.StatusText(resp.StatusCode), src)
	}

	//Open file handle
	out, err := os.Create(dst)
	defer out.Close()
	if err != nil {
		return err
	}

	//Copy response into file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return
}

//TODO: Multiple download sources based on Image Quality
//TODO: Base Image Quality off of a sample download from each site
