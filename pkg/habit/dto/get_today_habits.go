package dto

type GetTodayHabitsRequest struct{}

type GetTodayHabitsResponse struct {
	Habits []Habit `json:"habits"`
	// add today stats
}
