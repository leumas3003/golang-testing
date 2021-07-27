package test

import (
	"fmt"
	"golang-testing/api/utils/errors"
	mock "golang-testing/internal/mocks"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetCountryNotFound(t *testing.T) {
	fmt.Println("about to start test cases")

	//Create a mock with gomock
	controller := gomock.NewController(t)
	m := mock.NewMocklocationsServiceInterface(controller)
	defer controller.Finish()

	//Configure the mock
	m.EXPECT().GetCountry(gomock.Eq("AR")).Return(nil, &errors.ApiError{
		Status:  http.StatusNotFound,
		Message: "Country not found",
	})

	response, err := http.Get("http://localhost:3001/locations/countries/AR")
	fmt.Printf("Response, %v\n", response)
	fmt.Printf("Error, %v\n", err)
}
