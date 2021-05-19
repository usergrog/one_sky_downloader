package xml

import (
	"OneSkyDownloader/models"
	"OneSkyDownloader/utils"
	"encoding/xml"
	"io"
	"os"
	"strings"
)

func AndroidXMLWrite(localizations map[string]map[string]string) {
	filePath := "res/values"
	for lang, stringsMap := range localizations {

		err := os.MkdirAll(filePath+"-"+lang, 0700)
		utils.CheckError(err)
		fileOut, err := os.Create(filePath + "-" + lang + "/localizations.xml")
		utils.CheckError(err)
		_, err = fileOut.Write([]byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"))
		utils.CheckError(err)
		var rows []models.XMLString
		for _, stringKey := range utils.SortedMapKeys(stringsMap) {
			stringValue := stringsMap[stringKey]
			stringValue = strings.Replace(stringValue, "%@", "%s", -1)
			stringValue = strings.Replace(stringValue, "\n", " ", -1)
			row := models.XMLString{
				Name:      stringKey,
				Value:     stringValue,
				Formatted: !strings.Contains(stringValue, "%"),
			}
			rows = append(rows, row)
		}
		xmlStraps := models.XMLResources{
			XMLName:    xml.Name{Local: "resources"},
			XMLStrings: rows,
		}

		err = writeResourcesFile(fileOut, xmlStraps)
		utils.CheckError(err)

		err = fileOut.Close()
		utils.CheckError(err)
	}
}

func writeResourcesFile(writer io.Writer, xmlResources models.XMLResources) error {
	data, err := xml.MarshalIndent(xmlResources, "", "\t")
	if err != nil {
		return err
	}
	writer.Write(data)
	return nil
}
