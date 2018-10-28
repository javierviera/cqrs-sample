package main

import (
	"cqrs-test/dao/factory"
	"cqrs-test/dao/interfaces"
	"cqrs-test/utilities"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var dao interfaces.ProductDao

func main() {

	// Load configuration TODO Load for env
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create dao
	dao = factory.FactoryDao(config.Engine)

	e := echo.New()
	e.GET("/products/:id", getProduct)
	e.Logger.Fatal(e.Start(":1323"))

}

// e.GET("/products/:id", getProduct)
func getProduct(c echo.Context) error {
	// Product ID from path `products/:id`
	id := c.Param("id")
	product, err := dao.GetByID(id)
	if err != nil {
		c.Logger().Error(err)
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, product)
}
