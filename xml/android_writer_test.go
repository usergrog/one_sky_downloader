package xml

import (
	"OneSkyDownloader/parsers"
	"io/ioutil"
	"testing"
)

func TestAndroidXMLWrite(t *testing.T) {
	dat, err := ioutil.ReadFile("../raw.json")
	if err != nil {
		t.Fatalf(err.Error())
	}
	var rawJson = string(dat)
	var localizations = parsers.AndroidParse(rawJson)

	AndroidXMLWrite(localizations)

}
