package locations_provider

import (
	"encoding/json"
	"fmt"
	"golang-testing/api/domain/locations"
	"golang-testing/api/utils/errors"
	"io/ioutil"
	"net/http"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	response, err := http.Get(fmt.Sprintf(urlGetCountry, countryId))
	if err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient response when trying to get country %s", countryId),
		}
	}

	body, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode > 299 {
		var apiErr errors.ApiError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error interface when getting country %s", countryId),
			}
		}
		return nil, &apiErr
	}

	var result locations.Country
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryId),
		}
	}
	return &result, nil
}
