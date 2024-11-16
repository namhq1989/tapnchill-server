package dto

type GetStatsRequest struct {
	Date string `query:"date" validate:"required" message:"invalid_date"`
}

type GetStatsResponse struct {
	Stats []HabitDailyStats `json:"stats"`
}
