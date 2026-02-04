package service

import (
	"bug-tracker/internal/domain"
	"context"
)

type IssueRepository interface {
	Save(ctx context.Context, issue domain.Issue) error
	Update(ctx context.Context, issue domain.Issue) error
	Delete(ctx context.Context, id int) error
	ListIssue(ctx context.Context) ([]domain.Issue, error)
	FindByID(ctx context.Context, id int) (*domain.Issue, error)
}

type IssueService struct {
	repo IssueRepository
}

func NewIssueService(repo IssueRepository) *IssueService {
	return &IssueService{
		repo: repo,
	}
}

func (s *IssueService) CreateIssue(ctx context.Context, issue domain.Issue) error {
	if issue.Title == "" {
		return domain.ErrTitleRequired
	}
	return s.repo.Save(ctx, issue)
}

func (s *IssueService) UpdateIssue(ctx context.Context, issue domain.Issue) error {
	Issue, err := s.repo.FindByID(ctx, issue.ID)
	if err != nil {
		return err
	}

	if Issue == nil {
		return domain.ErrIssueNotFound
	}

	return s.repo.Update(ctx, issue)
}

func (s *IssueService) DeleteIssue(ctx context.Context, id int) error {
	Issue, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if Issue == nil {
		return domain.ErrIssueNotFound
	}

	return s.repo.Delete(ctx, id)
}

func (s *IssueService) GetAllIssues(ctx context.Context) ([]domain.Issue, error) {
	return s.repo.ListIssue(ctx)
}

func (s *IssueService) GetIssueById(ctx context.Context, id int) (*domain.Issue, error) {
	issue, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if issue == nil {
		return nil, domain.ErrIssueNotFound
	}

	return issue, nil
}
