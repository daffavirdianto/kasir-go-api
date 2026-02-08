package repositories

import (
	"database/sql"
	"fmt"
	"kasir-go-api/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	for i := range details {
		details[i].TransactionID = transactionID
		_, err = tx.Exec("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4)",
			transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

func (repo *TransactionRepository) ReportToday() (*models.Report, error) {
	report := models.Report{}

	// total revenue & total transaksi hari ini
	err := repo.db.QueryRow(`
		SELECT 
			COALESCE(SUM(total_amount), 0) AS total_revenue,
			COUNT(*)
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`).Scan(&report.TotalRevenue, &report.TotalTransactions)

	if err != nil {
		return nil, err
	}

	// produk terlaris hari ini
	err = repo.db.QueryRow(`
		SELECT 
			p.name,
			SUM(td.quantity) AS quantity_sold
		FROM transaction_details td
		JOIN transactions t ON t.id = td.transaction_id
		JOIN products p ON p.id = td.product_id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.name
		ORDER BY quantity_sold DESC
	`).Scan(&report.BestSellingProduct.Name, &report.BestSellingProduct.QuantitySold)

	// jika hari ini tidak ada transaksi
	if err == sql.ErrNoRows {
		report.BestSellingProduct = models.BestSeller{}
		return &report, nil
	}

	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (repo *TransactionRepository) ReportByDate(startDate string, endDate string) (*models.Report, error) {
	var report models.Report

	// total revenue & total transaksi
	err := repo.db.QueryRow(`
		SELECT 
			COALESCE(SUM(total_amount),0),
			COUNT(*)
		FROM transactions
		WHERE created_at::date BETWEEN $1 AND $2
	`, startDate, endDate).Scan(
		&report.TotalRevenue,
		&report.TotalTransactions,
	)
	if err != nil {
		return nil, err
	}

	// produk terlaris
	err = repo.db.QueryRow(`
		SELECT 
			p.name,
			COALESCE(SUM(td.quantity),0) as qty
		FROM transaction_details td
		JOIN products p ON p.id = td.product_id
		JOIN transactions t ON t.id = td.transaction_id
		WHERE t.created_at::date BETWEEN $1 AND $2
		GROUP BY p.name
		ORDER BY qty DESC
	`, startDate, endDate).Scan(
		&report.BestSellingProduct.Name,
		&report.BestSellingProduct.QuantitySold,
	)

	// jika belum ada transaksi
	if err == sql.ErrNoRows {
		report.BestSellingProduct = models.BestSeller{}
		return &report, nil
	}
	if err != nil {
		return nil, err
	}

	return &report, nil
}
