package domain

import "errors"

var (
	ErrTitleRequired       = errors.New("title is required")
	ErrDescriptionTooShort = errors.New("description too short")
	ErrIssueNotFound       = errors.New("issue not found")
)
