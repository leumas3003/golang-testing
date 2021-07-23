package controllers

import (
	"golang-testing/api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCountry(c echo.Context) error {
	country, err := services.LocationsService.GetCountry(c.Param("country_id"))
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(http.StatusOK, country)
}
