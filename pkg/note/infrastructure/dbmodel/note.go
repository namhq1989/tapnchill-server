package dbmodel

import (
	"time"

	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/note/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID          primitive.ObjectID `bson:"_id"`
	UserID      primitive.ObjectID `bson:"userId"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Data        *NoteData          `bson:"data"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}

func (n Note) ToDomain() domain.Note {
	note := domain.Note{
		ID:          n.ID.Hex(),
		UserID:      n.UserID.Hex(),
		Title:       n.Title,
		Description: n.Description,
		CreatedAt:   n.CreatedAt,
		UpdatedAt:   n.UpdatedAt,
	}

	if n.Data != nil {
		note.Data = &domain.NoteData{
			PageText:  n.Data.PageText,
			PageTitle: n.Data.PageTitle,
			PageURL:   n.Data.PageURL,
		}
	}

	return note
}

func (Note) FromDomain(note domain.Note) (*Note, error) {
	id, err := database.ObjectIDFromString(note.ID)
	if err != nil {
		return nil, apperrors.Common.InvalidNote
	}

	uid, err := database.ObjectIDFromString(note.UserID)
	if err != nil {
		return nil, apperrors.User.InvalidUserID
	}

	n := &Note{
		ID:          id,
		UserID:      uid,
		Title:       note.Title,
		Description: note.Description,
		CreatedAt:   note.CreatedAt,
		UpdatedAt:   note.UpdatedAt,
	}

	if note.Data != nil {
		n.Data = &NoteData{
			PageText:  note.Data.PageText,
			PageTitle: note.Data.PageTitle,
			PageURL:   note.Data.PageURL,
		}
	}

	return n, nil
}
