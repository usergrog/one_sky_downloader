package parsers

import (
	"OneSkyDownloader/models"
	"OneSkyDownloader/utils"
	"encoding/json"
)

func AndroidParse(jsonBody string) map[string]models.Translation {
	var parsedBody map[string]models.Translation
	err := json.Unmarshal([]byte(jsonBody), &parsedBody)
	utils.CheckError(err)
	return parsedBody
}
