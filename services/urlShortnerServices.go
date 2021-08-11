package services

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"
	"urlshortner/models"
	"urlshortner/utils"
)

func CreateShortURLService(request models.RequestData) (models.URLData, error) {
	url := models.URLData{}
	alias := utils.GenerateShortURL(request.URL)

	url, isFound := GetURLByAlias(alias)
	if isFound {
		return url, nil
	}
	url.Alias = alias
	url.OriginalURL = request.URL
	url.ShortURL = utils.BaseURL + alias
	url.CreatedOn = time.Now()

	if addURLurl(url) {
		return url, nil
	}

	return models.URLData{}, errors.New("ERROR_SHORT_URL_NOT_CREATED")

}

// This function used to get the all urls data from file.
func GetAllURLs() []models.URLData {
	urls := []models.URLData{}
	if _, err := os.Stat(utils.FileName); err == nil {

		// read file
		bytesArray, err := utils.ReadFile(utils.FileName)
		if err != nil {
			log.Println("GetAllURLs Error: ", err)
			return urls
		}

		err = json.Unmarshal(bytesArray, &urls)

		if err != nil {
			log.Println("GetAllURLs Error: ", err)
			return urls
		}

		return urls
	}
	return urls
}

// This function is used to add url into file
func addURLurl(url models.URLData) bool {
	// get all urls and append new url
	urls := append(GetAllURLs(), url)
	data, err := json.MarshalIndent(urls, "", " ")
	if err != nil {
		log.Println("addURLurl Error: ", err)
		return false
	}
	return utils.WriteFile(utils.FileName, data)
}

// get url by alias name
func GetURLByAlias(alias string) (models.URLData, bool) {
	urls := GetAllURLs()
	for _, url := range urls {
		if url.Alias == alias {
			return url, true
		}
	}
	return models.URLData{}, false
}
