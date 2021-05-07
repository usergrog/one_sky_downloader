package parsers

import (
	"io/ioutil"
	"testing"
)

func TestAndroidParse(t *testing.T) {
	dat, err := ioutil.ReadFile("../raw.json")
	if err != nil {
		t.Fatalf(err.Error())
	}
	var rawJson = string(dat)

	var localizations = AndroidParse(rawJson)
	for k,v := range localizations {
		t.Logf("%s", k)
		for k2, v2 := range v.Translation {
			t.Logf("%s - %s", k2, v2)
		}
	}
}
