package test

import (
	"fmt"
	"golang-testing/api/app"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	fmt.Println("about to start app")
	go app.StartServer()
	os.Exit(m.Run())
}
