package services

import (
	"time"
	"urlshortner/models"
	"urlshortner/utils"
)

func CreateShortURLService(request models.RequestData) (models.URLData, error) {
	url := models.URLData{}
	alias := utils.GenerateShortURL(request.URL)
	url.Alias = alias
	url.OriginalURL = request.URL
	url.ShortURL = utils.BaseURL + alias
	url.CreatedOn = time.Now()

	return url, nil

}
