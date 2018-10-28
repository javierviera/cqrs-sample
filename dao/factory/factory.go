package factory

import (
	"log"

	"cqrs-test/dao/interfaces"

	"cqrs-test/dao/mysql"
)

func FactoryDao(e string) interfaces.ProductDao {
	var i interfaces.ProductDao
	switch e {
	case "mysql":
		i = mysql.ProductImplMysql{}
	default:
		log.Fatalf("Database engine %s not implemented", e)
		return nil
	}

	return i
}
