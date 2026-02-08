package models

type Report struct {
	TotalRevenue       int        `json:"total_revenue"`
	TotalTransactions  int        `json:"total_transactions"`
	BestSellingProduct BestSeller `json:"best_selling_product"`
}

type BestSeller struct {
	Name         string `json:"name"`
	QuantitySold int    `json:"quantity_sold"`
}
