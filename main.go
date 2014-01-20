package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path"
)

const (
	octgnGameId string = "0f38e453-26df-4c04-9d67-6d43de939c77"
	markerSetId string = "21bf9e05-fb23-4b1d-b89a-398f671f5999"
)

//Get important OCTGN Netrunner directory paths
func getPaths() (installPath, setPath, imgPath string) {
	curUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	installPath = path.Join(curUser.HomeDir, "Documents", "OCTGN")
	setPath = path.Join(installPath, "GameDatabase", octgnGameId, "Sets")
	imgPath = path.Join(installPath, "ImageDatabase", octgnGameId, "Sets")
	return
}


type Task struct{
	Target string
	Card CardInfo
}


func main() {
	installPath, setPath, imgPath := getPaths()
	_ = installPath
	_ = setPath
	_ = imgPath
	
	//Foreach set
	setDirs, err := ioutil.ReadDir(setPath)
	if err != nil {
		log.Fatal(err)
	}
	workChan := make(chan Task)
	for _, curSet := range setDirs {
		if curSet.Name() == markerSetId{
			continue
		}		
		setColl := parseSetXML(fmt.Sprintf("%s/%s/set.xml", setPath, curSet.Name()))
		_ = setColl
		
		for _,curCard:= range setColl{
			targetFile := path.Join(imgPath, curCard.SetID, "Cards", curCard.ID,".png")
			//Check card exists, set quality
			
			
			workChan <- Task{ targetFile, curCard}
		}		
	}
	
	/*
	
def downloadSet(setName, setID, cardSet chan cardWork):
    # Setup dl directory
    targetDir = r"{}\{}\Cards".format(imagePath, setID)
    if not os.path.isdir(targetDir):
        os.makedirs(targetDir)

    for curCard in cardSet:
        curID = curCard.attrib["id"]
        curName = curCard.attrib["name"]
        targetFile = r"{}\{}.png".format(targetDir, curID)

        # First attempt from netrunnercards.info
        print("Downloading {} - {}".format(setName, curName))
        targetURL = "{}{}.png".format(priAssetURL, curID[-5:])
        if not downloadImage(targetURL, targetFile):
            # Second attempt from cardgamedb
            print("\tAttempting alternative download")
            targetURL = "{}ffg_{}-{}.png".format(secAssetURL, curName.lower().replace(' ', '-'), setName.lower().replace(' ', '-'))
            downloadImage(targetURL, targetFile)

	*/
	
	//foreach dir in setPath:
	//format set xml path
	//get setInfo from xml
	//if curSet not netMarkerSet:
	//download set
}




/* 

Channel of CardInfo & targetPath



FUNCTION targetPath use sources
	Check if targetPath exists
	if exists, get resolution
	while resolution of next source > image try:
		download and save image
		if success break
	









/*
func main() {
		v := new(set)
		err := xml.Unmarshal([]byte(xmlString ),&v)
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Println(v.Name)
		for _,cur := range v.Cards.Cards{
			fmt.Println(cur)
		}
	fmt.Println(getPaths())
}
*/
/*
def getCardSet(targetSetXML):
    tree = xml.etree.ElementTree.parse(targetSetXML)
    return (tree.getroot().attrib['name'], tree.findall(".//card"))
*/




/*


def getCardSet(targetSetXML):
    tree = xml.etree.ElementTree.parse(targetSetXML)
    return (tree.getroot().attrib['name'], tree.findall(".//card"))


def downloadImage(targetURL, outFilePath):
    response = None
    try:
        response = urllib2.urlopen(targetURL)
        with open(outFilePath, "wb") as outFile:
            outFile.write(response.read())
    except Exception as err:
        print("\tError in getting file from {}:\n\t\t{}".format(targetURL, err))
        return False
    finally:
        try:
            response.close()
        except:
            pass
    return True


def downloadSet(setName, setID, cardSet):
    # Setup dl directory
    targetDir = r"{}\{}\Cards".format(imagePath, setID)
    if not os.path.isdir(targetDir):
        os.makedirs(targetDir)

    for curCard in cardSet:
        curID = curCard.attrib["id"]
        curName = curCard.attrib["name"]
        targetFile = r"{}\{}.png".format(targetDir, curID)

        # First attempt from netrunnercards.info
        print("Downloading {} - {}".format(setName, curName))
        targetURL = "{}{}.png".format(priAssetURL, curID[-5:])
        if not downloadImage(targetURL, targetFile):
            # Second attempt from cardgamedb
            print("\tAttempting alternative download")
            targetURL = "{}ffg_{}-{}.png".format(secAssetURL, curName.lower().replace(' ', '-'), setName.lower().replace(' ', '-'))
            downloadImage(targetURL, targetFile)


*/
