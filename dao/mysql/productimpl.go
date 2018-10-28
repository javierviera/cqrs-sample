package mysql

import (
	"cqrs-test/model"
)

// ProductImplMysql struct
type ProductImplMysql struct {
}

// Create implementation
func (dao ProductImplMysql) Create(p *model.Product) error {
	query := "INSERT INTO products (name, description) VALUES(?, ?)"
	db := get()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(p.Name, p.Description)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = string(id)
	return nil
}

// GetByID implementation
func (dao ProductImplMysql) GetByID(i string) (model.Product, error) {
	query := "SELECT id, name, description FROM products where id = ?"
	product := model.Product{}
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return product, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(i)
	if err != nil {
		return product, err
	}

	for rows.Next() {
		var row model.Product
		err := rows.Scan(&row.ID, &row.Name, &row.Description)
		if err != nil {
			return model.Product{}, err
		}
		return row, nil
	}

	return product, nil

}

// GetAll function
func (dao ProductImplMysql) GetAll() ([]model.Product, error) {
	query := "SELECT id, name, description, email FROM products"
	products := make([]model.Product, 0)
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return products, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return products, err
	}

	for rows.Next() {
		var row model.Product
		err := rows.Scan(&row.ID, &row.Name, &row.Description)
		if err != nil {
			return nil, err
		}

		products = append(products, row)
	}

	return products, nil

}
