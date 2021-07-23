package controllers

import (
	"encoding/json"
	"golang-testing/api/domain/locations"
	"golang-testing/api/services"
	"golang-testing/api/utils/errors"
	mock "golang-testing/internal/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	getCountryFunc func(countryId string) (*locations.Country, *errors.ApiError)
)

type locationsServiceMock struct{}

func (*locationsServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return getCountryFunc(countryId)
}

func TestGetCountryNotFound(t *testing.T) {
	// Mock LocationsService methods:
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return nil, &errors.ApiError{Status: http.StatusNotFound, Message: "Country not found"}
	}
	services.LocationsService = &locationsServiceMock{}

	//Create Context for the request with echo ("MOCK")
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/locations/countries/:country_id")
	c.SetParamNames("country_id")
	c.SetParamValues("AR")
	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, rec.Code)

	var apiErr errors.ApiError
	err := json.Unmarshal(rec.Body.Bytes(), &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}

func TestGetCountryNoError(t *testing.T) {
	// Mock LocationsService methods:
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return &locations.Country{Id: "AR", Name: "Argentina"}, nil
	}
	services.LocationsService = &locationsServiceMock{}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/locations/countries/:country_id")
	c.SetParamNames("country_id")
	c.SetParamValues("AR")
	GetCountry(c)

	assert.EqualValues(t, http.StatusOK, rec.Code)

	var country locations.Country
	err := json.Unmarshal(rec.Body.Bytes(), &country)
	assert.Nil(t, err)

	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}

func TestGetCountryNotFoundMock(t *testing.T) {
	//Create a mock with gomock
	controller := gomock.NewController(t)
	m := mock.NewMocklocationsServiceInterface(controller)
	defer controller.Finish()

	//Create echo Context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/locations/countries/:country_id")

	//Configure the mock
	m.EXPECT().GetCountry(gomock.Eq("AR")).Return(nil, &errors.ApiError{
		Status:  http.StatusNotFound,
		Message: "Country not found",
	})

	_, err := m.GetCountry("AR")

	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)

}

func TestGetCountryNotErrordMock(t *testing.T) {
	//Create a mock with gomock
	controller := gomock.NewController(t)
	m := mock.NewMocklocationsServiceInterface(controller)
	defer controller.Finish()

	//Create echo Context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/locations/countries/:country_id")

	//Configure the mock
	m.EXPECT().GetCountry(gomock.Eq("AR")).Return(&locations.Country{
		Id:   "AR",
		Name: "Argentina",
	}, nil)

	country, _ := m.GetCountry("AR")

	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)

}
