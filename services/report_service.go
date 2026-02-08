package services

import (
	"kasir-go-api/models"
	"kasir-go-api/repositories"
)

type ReportService struct {
	repo *repositories.TransactionRepository
}

func NewReportService(repo *repositories.TransactionRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetReportToday() (*models.Report, error) {
	return s.repo.ReportToday()
}

func (s *ReportService) ReportByDate(startDate string, endDate string) (*models.Report, error) {
	return s.repo.ReportByDate(startDate, endDate)
}
