package handler

import (
	"bug-tracker/internal/domain"
	"bug-tracker/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type IssueHandler struct {
	service *service.IssueService
}

func NewIssueHandler(service *service.IssueService) *IssueHandler {
	return &IssueHandler{
		service: service,
	}
}

type CreateIssueRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

func (h *IssueHandler) CreateIssue(w http.ResponseWriter, r *http.Request) {
	var req CreateIssueRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	issue := domain.Issue{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
	}

	ctx := r.Context()
	if err := h.service.CreateIssue(ctx, issue); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "issue created",
	})
}

type IssueResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

func (h *IssueHandler) GetAllIssues(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	issues, err := h.service.GetAllIssues(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var resp []IssueResponse
	for _, issue := range issues {
		resp = append(resp, IssueResponse{
			ID:          issue.ID,
			Title:       issue.Title,
			Description: issue.Description,
			Priority:    issue.Priority,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

type UpdateIssueRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

func (h *IssueHandler) UpdateIssue(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/issues/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Issue ID", http.StatusBadRequest)
		return
	}

	var req UpdateIssueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	issue := domain.Issue{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
	}

	ctx := r.Context()
	if err := h.service.UpdateIssue(ctx, issue); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "issue updated",
	})
}
func (h *IssueHandler) DeleteIssue(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/issues/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Issue ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err = h.service.DeleteIssue(ctx, id)
	if err != nil {
		if err == domain.ErrIssueNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *IssueHandler) GetIssueByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/issues/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Issue ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	issue, err := h.service.GetIssueById(ctx, id)
	if err != nil {
		if err == domain.ErrIssueNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	resp := IssueResponse{
		ID:          issue.ID,
		Title:       issue.Title,
		Description: issue.Description,
		Priority:    issue.Priority,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
