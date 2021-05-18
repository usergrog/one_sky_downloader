package main

import (
	"OneSkyDownloader/parsers"
	"OneSkyDownloader/utils"
	"OneSkyDownloader/xml"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	log.Println("Start OneSky downloader")

	config := utils.ReadConfig()
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	data := []byte(timestamp + config.ApiSecret)
	devHash := md5.Sum(data) //md5(concatenate(<timestamp>, <api_secret>))

	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://platform.api.onesky.io/1/projects/"+config.ProjectId+"/translations/multilingual", nil)

	q := req.URL.Query()
	q.Add("api_key", config.ApiKey)
	q.Add("timestamp", timestamp)
	q.Add("dev_hash", hex.EncodeToString(devHash[:]))
	q.Add("source_file_name", config.FileName)
	q.Add("file_format", "I18NEXT_MULTILINGUAL_JSON")

	req.URL.RawQuery = q.Encode()
	resp, _ := client.Do(req)

	log.Println(resp.Status)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		utils.CheckError(err)
	}(resp.Body)
	respBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("Parse response")
	parsedBody := parsers.Parse(string(respBody))

	log.Println("Write results")
	if config.Format == "android" {
		xml.AndroidXMLWrite(parsedBody)
	} else {
		xml.IOSWrite(parsedBody)
	}
	log.Println("Finished")
}
