package parsers

import (
	"io/ioutil"
	"testing"
)

func TestParseOneSky(t *testing.T) {
	dat, err := ioutil.ReadFile("../raw.json")
	if err != nil {
		t.Fatalf(err.Error())
	}
	var rawJson = string(dat)

	var localizations = ParseOneSky(rawJson)
	for k, v := range localizations {
		t.Logf("%s", k)
		for k2, v2 := range v {
			t.Logf("%s - %s", k2, v2)
		}
	}
}
