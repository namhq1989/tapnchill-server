package dto

type GetHabitsRequest struct{}

type GetHabitsResponse struct {
	Habits []Habit `json:"habits"`
}
