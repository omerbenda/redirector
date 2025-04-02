package db

import (
	"encoding/json"
	"log"
	"os"
)

func readUrlFile(fileName string) map[string]string {
	fileUrlMap := make(map[string]string)
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		err = os.WriteFile(fileName, []byte("{}"), 0644)

		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}

	fileContent, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(fileContent, &fileUrlMap)

	if err != nil {
		log.Fatal(err)
	}

	return fileUrlMap
}

func WriteMapToFile(fileName string, urlHashMap map[string]string) {
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(urlHashMap)

	if err != nil {
		log.Fatal(err)
	}
}
