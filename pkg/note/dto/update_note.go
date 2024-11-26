package dto

type UpdateNoteRequest struct {
	Title       string    `json:"title" validate:"required" message:"invalid_name"`
	Description string    `json:"description"`
	Data        *NoteData `json:"data"`
}

type UpdateNoteResponse struct {
	ID string `json:"id"`
}
