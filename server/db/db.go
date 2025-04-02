package db

import (
	"crypto/sha256"
	"encoding/hex"
)

const DB_FILE_NAME = "./urls.json"

var hashToUrlMap map[string]string

func Read() {
	hashToUrlMap = readUrlFile(DB_FILE_NAME)
}

func GetValue(hash string) (string, bool) {
	url, ok := hashToUrlMap[hash]

	return url, ok
}

func SetValue(url string) string {
	hash := GenerateHash(url)

	hashToUrlMap[hash] = url

	WriteMapToFile(DB_FILE_NAME, hashToUrlMap)

	return hash
}

func GenerateHash(url string) string {
	hasher := sha256.New()

	hasher.Write([]byte(url))
	hashBytes := hasher.Sum(nil)

	return hex.EncodeToString(hashBytes)
}
