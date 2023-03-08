package repository

import (
	"sirawit/shop/internal/config"
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/random"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var testProductDB *gorm.DB
var testPruductQuery ProductQuery

func TestConnectToProductDB(t *testing.T) {
	var err error
	config, err := config.LoadProductConfig("../../cmd/product")
	assert.NoError(t, err)
	testProductDB, err = ConnectToProductDB(config.DSN)
	assert.NoError(t, err)
	testPruductQuery = NewProductQuery(testProductDB)

}

func TestCreateProduct(t *testing.T) {
	qty := 3
	for i := 0; i < qty; i++ {
		product := model.Product{
			Name:     random.RandomUsername(),
			Price:    float64(random.RandomUInt64(0, 1000)),
			Details:  random.RandomString(20),
			ImageUrl: random.RandomString(20),
		}
		result, err := testPruductQuery.CreateProduct(product)
		assert.NoError(t, err)
		assert.Equal(t, result.Details, product.Details)
		assert.Equal(t, result.Price, product.Price)
		assert.Equal(t, result.ImageUrl, product.ImageUrl)
		assert.Equal(t, result.Name, product.Name)
		assert.NotNil(t, result.ID)
	}
}
