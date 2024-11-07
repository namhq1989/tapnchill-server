package dto

import "time"

type Goal struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stats       GoalStats `json:"stats"`
	IsCompleted bool      `json:"isCompleted"`
	CreatedAt   time.Time `json:"createdAt"`
}

type GoalStats struct {
	TotalTask          int `json:"totalTask"`
	TotalCompletedTask int `json:"totalCompletedTask"`
}
