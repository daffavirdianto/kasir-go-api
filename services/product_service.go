package services

import (
	"errors"
	"kasir-go-api/models"
	"kasir-go-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]models.ProductResponse, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var response []models.ProductResponse
	for _, d := range data {
		response = append(response, models.ProductResponse{
			ID:       d.ID,
			Name:     d.Name,
			Price:    d.Price,
			Stock:    d.Stock,
			Category: d.CategoryName,
		})
	}

	return response, nil
}

func (s *ProductService) Create(data *models.Product) error {
	return s.repo.Create(data)
}

func (s *ProductService) GetByID(id int) (*models.ProductResponse, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("produk tidak ditemukan")
	}

	return &models.ProductResponse{
		ID:       data[0].ID,
		Name:     data[0].Name,
		Price:    data[0].Price,
		Stock:    data[0].Stock,
		Category: data[0].CategoryName,
	}, nil
}

func (s *ProductService) Update(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
