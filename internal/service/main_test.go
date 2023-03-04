package service

import (
	"log"
	"os"
	"sirawit/shop/internal/config"
	"sirawit/shop/mock"
	"testing"
)

var testUserService UserService

func TestMain(m *testing.M) {
	var err error
	config, err := config.LoadUserConfig("../../cmd/user")
	if err != nil {
		log.Fatal("cannot load config file", err)
	}
	testDB := mock.NewUserRepositoryMock()
	testUserService = NewUserService(testDB, config)
	os.Exit(m.Run())
}
