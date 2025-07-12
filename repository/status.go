package repository

type Status string

const (
	StatusCreated   Status = "CREATED"
	StatusModified  Status = "MODIFIED"
	StatusCompleted Status = "COMPLETED"
)
