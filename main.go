package main

import (
	"fmt"
	"os/user"
	"path"
	"strings"
	"bytes"
	"log"
)

const(	
	octgnGameId string = "0f38e453-26df-4c04-9d67-6d43de939c77"
	markerSetId string = "21bf9e05-fb23-4b1d-b89a-398f671f5999"
)

type CardInfo struct{
	ID string
	Set string
	Name string
}
type CustomURL func(info CardInfo) string
type AssetSource struct{
	Quality uint
	ComposeURL CustomURL
}
var assetList []AssetSource = []AssetSource{
	{
		Quality: 125400,
		ComposeURL: func(info CardInfo) string{
			startIdx := len(info.ID)-5
			buffer := bytes.NewBufferString("http://netrunnerdb.com/web/bundles/netrunnerdbcards/images/cards/en/")
			buffer.WriteString(info.ID[startIdx:])
			buffer.WriteString(".png")
			return buffer.String()
		},
	},
	{
		Quality: 58870,
		ComposeURL: func(info CardInfo) string{			
			buffer := bytes.NewBufferString("http://www.cardgamedb.com/forums/uploads/an/")
			buffer.WriteString("ffg_")
			buffer.WriteString(strings.Replace(strings.ToLower(info.Name)," ","-",-1))
			buffer.WriteString("-")
			buffer.WriteString(strings.Replace(strings.ToLower(info.Set)," ","-",-1))
			buffer.WriteString(".png")
			return buffer.String()
		},
	},
}
	
func getPaths() (installPath , setPath , imgPath string){
	curUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	installPath = path.Join(curUser.HomeDir, "Documents", "OCTGN")
	setPath = path.Join(installPath, "GameDatabase", octgnGameId, "Sets")
	imgPath = path.Join(installPath, "ImageDatabase", octgnGameId, "Sets")
	return
}

func main() {
	
	installPath, setPath, imgPath := getPaths()

	#foreach dir in setPath:
		#format set xml path
		#get setInfo from xml
		#if curSet not netMarkerSet:
			#download set
	
	
	
	
	tmp := CardInfo{
		ID:"78923478295",
		Set:"Set Name Here",
		Name:"Card Name",
		}
	fmt.Println(assetList[1].ComposeURL(tmp))
}





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