package xml

import (
	"OneSkyDownloader/models"
	"OneSkyDownloader/utils"
	"encoding/xml"
	"io"
	"os"
	"strings"
)

func AndroidXMLWrite(localizations map[string]models.Translation) {
	filePath := "res/values"
	for lang, a := range localizations {

		err := os.MkdirAll(filePath+"-"+lang, 0700)
		utils.CheckError(err)
		fileOut, err := os.Create(filePath + "-" + lang + "/localizations.xml")
		utils.CheckError(err)
		_, err = fileOut.Write([]byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"))
		utils.CheckError(err)
		var rows []models.XMLString
		for key, v := range a.Translation {
			if strings.Contains(v, "%@") {
				v = strings.Replace(v, "%@", "%s", -1)
			}
			row := models.XMLString{
				Name:      key,
				Value:     v,
				Formatted: !strings.Contains(v, "%"),
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