package dto

type CreateFeedbackRequest struct {
	Email    string `json:"email"`
	Feedback string `json:"feedback" validate:"required" message:"invalid_feedback"`
}

type CreateFeedbackResponse struct {
	ID string `json:"id"`
}
