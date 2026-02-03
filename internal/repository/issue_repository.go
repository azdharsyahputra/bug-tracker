package repository

import (
	"bug-tracker/internal/domain"
	"context"
	"database/sql"
)

type IssueRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *IssueRepository {
	return &IssueRepository{db: db}
}

func (r *IssueRepository) Save(ctx context.Context, issue domain.Issue) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO issues (title, description, priority) VALUES(?, ?, ?)",
		issue.Title,
		issue.Description,
		issue.Priority)

	return err
}

func (r *IssueRepository) Update(ctx context.Context, issue domain.Issue) error {
	_, err := r.db.ExecContext(ctx, "UPDATE issues SET title = ? description = ? priority = ? WHERE id = ?",
		issue.Title,
		issue.Description,
		issue.Priority)

	return err
}

func (r *IssueRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM issues WHERE id = ?", id)
	return err
}

func (r *IssueRepository) ListIssue(ctx context.Context) ([]domain.Issue, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, description, priority FROM issues")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var issues []domain.Issue

	for rows.Next() {
		var issue domain.Issue
		if err := rows.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Priority); err != nil {
			return nil, err
		}

		issues = append(issues, issue)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return issues, nil
}

func (r *IssueRepository) FindByID(ctx context.Context, id int) (*domain.Issue, error) {
	var issue domain.Issue

	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, title, description, priority FROM issues WHERE id = ?",
		id,
	).Scan(
		&issue.ID,
		&issue.Title,
		&issue.Description,
		&issue.Priority,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &issue, nil
}
