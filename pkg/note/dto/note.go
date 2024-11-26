package dto

import (
	"github.com/namhq1989/tapnchill-server/internal/utils/httprespond"
	"github.com/namhq1989/tapnchill-server/pkg/note/domain"
)

type Note struct {
	ID          string                    `json:"id"`
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Data        *NoteData                 `json:"data"`
	CreatedAt   *httprespond.TimeResponse `json:"createdAt"`
	UpdatedAt   *httprespond.TimeResponse `json:"updatedAt"`
}

func (Note) FromDomain(note domain.Note) Note {
	n := Note{
		ID:          note.ID,
		Title:       note.Title,
		Description: note.Description,
		CreatedAt:   httprespond.NewTimeResponse(note.CreatedAt),
		UpdatedAt:   httprespond.NewTimeResponse(note.UpdatedAt),
	}

	if note.Data != nil {
		n.Data = &NoteData{
			PageText:  note.Data.PageText,
			PageTitle: note.Data.PageTitle,
			PageURL:   note.Data.PageURL,
		}
	}

	return n
}
