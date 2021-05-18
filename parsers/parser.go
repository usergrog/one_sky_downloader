package parsers

import (
	"OneSkyDownloader/models"
	"OneSkyDownloader/utils"
	"encoding/json"
	"strings"
)

func ParseOneSky(jsonBody string) map[string]map[string]string {
	var parsedBody map[string]models.Translation
	err := json.Unmarshal([]byte(jsonBody), &parsedBody)
	utils.CheckError(err)

	responseBody := make(map[string]map[string]string)

	for langKey, stringsArray := range parsedBody {
		langEntries := make(map[string]string)
		m := stringsArray.Translation.(map[string]interface{})
		for stringKey, stringValue := range m {
			switch stringValue.(type) {
			case string:
				langEntries[stringKey] = strings.TrimSpace(stringValue.(string))
			case []interface{}:
				var joinResult string
				for _, u := range stringValue.([]interface{}) {
					if len(joinResult) > 0 {
						joinResult = joinResult + "\n\r"
					}
					joinResult = joinResult + u.(string)
				}
				langEntries[stringKey] = strings.TrimSpace(joinResult)
			}
		}
		responseBody[langKey] = langEntries
	}

	return responseBody
}
