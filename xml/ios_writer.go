package xml

import (
	"OneSkyDownloader/utils"
	"os"
	"strings"
)

func IOSWrite(localizations map[string]map[string]string) {
	filePath := "Localization"
	for lang, stringsMap := range localizations {

		err := os.MkdirAll(filePath+"/"+lang+".lproj", 0700)
		utils.CheckError(err)
		fileOut, err := os.Create(filePath + "/" + lang + ".lproj" + "/Localizable.strings")
		utils.CheckError(err)
		for _, stringKey := range utils.SortedMapKeys(stringsMap) {
			stringValue := stringsMap[stringKey]
			//v = strings.Replace(v, "\n", " ", -1)
			stringValue = strings.Replace(stringValue, `"`, `\"`, -1)
			data := "\"" + stringKey + "\"" + " = \"" + stringValue + "\";\n"
			_, err := fileOut.Write([]byte(data))
			utils.CheckError(err)
		}

		err = fileOut.Close()
		utils.CheckError(err)
	}
}
