package repository

import (
	"log"
	"os"
	"sirawit/shop/pkg/config"
	"testing"

	"gorm.io/gorm"
)

var testDB *gorm.DB
var testUserQuery UserQuery

func TestMain(m *testing.M) {
	var err error
	config, err := config.LoadUserConfig("../../cmd/user")
	if err != nil {
		log.Fatal("cannot load config file", err)
	}
	testDB, err = ConnectToUserDB(config.DSN)
	if err != nil {
		log.Fatal("cannot connect to user db", err)
	}
	testUserQuery = NewUserRepository(testDB)

	os.Exit(m.Run())
}
