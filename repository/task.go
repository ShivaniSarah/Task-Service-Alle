package repository

type Task struct {
	ID     	*uint   `gorm:"primaryKey" json:"id"`
	Title  	*string `json:"title"`
	Description *string `json:"description"`
	Status 	*string `json:"status"`
}
