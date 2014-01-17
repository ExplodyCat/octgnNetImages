'''
Created on Dec 16, 2013

@author: Jerry
'''
import os
import sys
import xml.etree.ElementTree
import urllib2


octgnGameId = "0f38e453-26df-4c04-9d67-6d43de939c77"
netMarkerSet = "21bf9e05-fb23-4b1d-b89a-398f671f5999"
installPath = os.path.expanduser(r"~\Documents\OCTGN")
setPath = os.path.join(installPath, r"GameDatabase\{}\Sets".format(octgnGameId))
imagePath = os.path.join(installPath, r"ImageDatabase\{}\Sets".format(octgnGameId))


priAssetURL = "http://netrunnerdb.com/web/bundles/netrunnerdbcards/images/cards/en/"
secAssetURL = "http://www.cardgamedb.com/forums/uploads/an/"


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


def cmdLineMain():
    for curSet in os.listdir(setPath):
        setName, cardSet = getCardSet("{}\{}\set.xml".format(setPath, curSet))
        if curSet != netMarkerSet:
            downloadSet(setName, curSet, cardSet)
    print "---Application completed---"
    return 0

if __name__ == '__main__':
    sys.exit(cmdLineMain())
