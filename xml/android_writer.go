package xml

import (
	"OneSkyDownloader/utils"
	"html"
	"os"
	"strings"
)

func AndroidXMLWrite(localizations map[string]map[string]string) {
	filePath := "res/values"
	for lang, stringsMap := range localizations {
		runes := []rune(lang)
		lang = string(runes[0:2])
		if lang != "en" {
			lang = "-" + lang
		} else {
			lang = ""
		}
		err := os.MkdirAll(filePath+lang, 0700)
		utils.CheckError(err)
		fileOut, err := os.Create(filePath + lang + "/localizable.xml")
		utils.CheckError(err)
		_, err = fileOut.Write([]byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"))
		utils.CheckError(err)
		_, err = fileOut.Write([]byte("<resources>\n"))
		utils.CheckError(err)
		//var rows []models.XMLString
		for _, stringKey := range utils.SortedMapKeys(stringsMap) {
			stringValue := html.UnescapeString(stringsMap[stringKey])
			stringValue = strings.Replace(stringValue, "%@", "%s", -1)
			stringValue = strings.Replace(stringValue, "\n", " ", -1)
			//stringValue = strings.Replace(stringValue, "<", "'", -1)
			//stringValue = strings.Replace(stringValue, ">", "'", -1)
			stringValue = strings.Replace(stringValue, `"`, `\"`, -1)
			stringValue = strings.Replace(stringValue, "'", "\\'", -1)
			var line = "\t<string name=\"" + strings.TrimSpace(stringKey) + "\" "
			if !strings.Contains(stringValue, "%") {
				line = line + "formatted=\"true\""
			}
			line = line + ">" + html.EscapeString(stringValue) + "</string>\n"
			_, err = fileOut.Write([]byte(line))
			utils.CheckError(err)
		}

		_, err = fileOut.Write([]byte("</resources>\n"))
		utils.CheckError(err)

		err = fileOut.Close()
		utils.CheckError(err)
	}
}
