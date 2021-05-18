package xml

import (
	"OneSkyDownloader/utils"
	"os"
)

func IOSWrite(localizations map[string]map[string]string) {
	filePath := "Localization"
	for lang, a := range localizations {

		err := os.MkdirAll(filePath+"/"+lang+".lproj", 0700)
		utils.CheckError(err)
		fileOut, err := os.Create(filePath + "/" + lang + ".lproj" + "/Localizable.strings")
		utils.CheckError(err)
		for key, v := range a {
			//v = strings.Replace(v, "\n", " ", -1)
			//v = strings.Replace(v, "\r", " ", -1)
			//"STR_TRANSACTION_ID" = "Transaction ID";
			data := "\"" + key + "\"" + " = \"" + v + "\";\n"
			_, err := fileOut.Write([]byte(data))
			utils.CheckError(err)
		}

		err = fileOut.Close()
		utils.CheckError(err)
	}
}
