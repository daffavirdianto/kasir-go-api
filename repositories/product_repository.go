package repositories

import (
	"database/sql"
	"errors"
	"kasir-go-api/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

type ProductWithCategory struct {
	models.Product
	CategoryName string
}

func (repo *ProductRepository) GetAll() ([]ProductWithCategory, error) {
	query := "SELECT p.id, p.name, p.price, p.stock, p.category_id, c.name FROM products p JOIN categories c ON p.category_id = c.id"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []ProductWithCategory
	for rows.Next() {
		var r ProductWithCategory
		err := rows.Scan(&r.ID, &r.Name, &r.Price, &r.Stock, &r.CategoryID, &r.CategoryName)
		if err != nil {
			return nil, err
		}
		results = append(results, r)
	}

	return results, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryID).Scan(&product.ID)
	return err
}

// GetByID - ambil produk by ID
func (repo *ProductRepository) GetByID(id int) ([]ProductWithCategory, error) {
	query := "SELECT p.id, p.name, p.price, p.stock, p.category_id, c.name FROM products p JOIN categories c ON p.category_id = c.id WHERE p.id = $1"

	var p ProductWithCategory
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID, &p.CategoryName)
	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	return []ProductWithCategory{p}, nil
}

func (repo *ProductRepository) Update(product *models.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoryID, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return err
}
