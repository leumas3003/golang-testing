package controllers

import (
	"golang-testing/api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetCountry godoc
// @Summary GetCountry
// @Description GetCoundry by countryID (external api)
// @Accept json
// @Produce json
// @Success 201 {object} locations.Country
// @Failure 404 {object} errors.ApiError
// @Failure 500 {object} errors.ApiError
// @Failure 400 {object} errors.ApiError
// @Param country_id path string true "Country ID"
// @Router /locations/countries/{country_id} [get]
func GetCountry(c echo.Context) error {
	country, err := services.LocationsService.GetCountry(c.Param("country_id"))
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(http.StatusOK, country)
}
