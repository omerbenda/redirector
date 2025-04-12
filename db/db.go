package db

import (
	"github.com/omerbenda/redirector/id"
)

const DB_FILE_NAME = "./urls.json"
const ID_LENGTH = 25

var UrlIdMap map[string]string

func Read() {
	UrlIdMap = readUrlFile(DB_FILE_NAME)
}

func GetValue(urlId string) (string, bool) {
	url, ok := UrlIdMap[urlId]

	return url, ok
}

func SetValue(url string) string {
	urlId := id.GenerateId(ID_LENGTH)

	for _, ok := UrlIdMap[urlId]; ok; {
		urlId = id.GenerateId(ID_LENGTH)
	}

	UrlIdMap[urlId] = url

	WriteMapToFile(DB_FILE_NAME, UrlIdMap)

	return urlId
}

func UpdateValue(id string, url string) bool {
	_, exists := UrlIdMap[id]

	if exists {
		UrlIdMap[id] = url
		WriteMapToFile(DB_FILE_NAME, UrlIdMap)
	}

	return exists
}

func GetCount() int {
	return len(UrlIdMap)
}
