package dto

type CreateNoteRequest struct {
	Title       string    `json:"title" validate:"required" message:"invalid_name"`
	Description string    `json:"description"`
	Data        *NoteData `json:"data"`
}

type CreateNoteResponse struct {
	ID string `json:"id"`
}
