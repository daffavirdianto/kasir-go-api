package handlers

import (
	"encoding/json"
	"kasir-go-api/services"
	"net/http"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate != "" && endDate != "" {
		h.GetReportByDate(w, r, startDate, endDate)
		return
	}

	h.GetReportToday(w, r)
}

// GET /api/report/hari-ini
func (h *ReportHandler) GetReportToday(w http.ResponseWriter, r *http.Request) {
	report, err := h.service.GetReportToday()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

// Get api/report?start_date=2026-01-01&end_date=2026-02-01
func (h *ReportHandler) GetReportByDate(w http.ResponseWriter, r *http.Request, startDate string, endDate string) {
	reports, err := h.service.ReportByDate(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}
