package domain

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/database"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/internal/utils/manipulation"
)

type NoteRepository interface {
	Create(ctx *appcontext.AppContext, note Note) error
	Update(ctx *appcontext.AppContext, note Note) error
	Delete(ctx *appcontext.AppContext, noteID string) error
	CountByUserID(ctx *appcontext.AppContext, userID string) (int64, error)
	FindByID(ctx *appcontext.AppContext, noteID string) (*Note, error)
	Sync(ctx *appcontext.AppContext, userID string, updatedAt time.Time, numOfNotes int64) ([]Note, error)
}

const NumOfNotesEachSync int64 = 20

type Note struct {
	ID          string
	UserID      string
	Title       string
	Description string
	Data        *NoteData
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewNote(userID string, title string, description string, data *NoteData) (*Note, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	note := &Note{
		ID:        database.NewStringID(),
		UserID:    userID,
		Data:      data,
		CreatedAt: manipulation.NowUTC(),
		UpdatedAt: manipulation.NowUTC(),
	}

	if err := note.SetTitle(title); err != nil {
		return nil, err
	}
	if err := note.SetDescription(description); err != nil {
		return nil, err
	}

	return note, nil
}

func (n *Note) SetTitle(title string) error {
	if len(title) < 3 || len(title) > 50 {
		return apperrors.Common.InvalidName
	}

	n.Title = title
	n.SetUpdatedAt()
	return nil
}

func (n *Note) SetDescription(description string) error {
	if len(description) > 300 {
		return apperrors.Common.InvalidDescription
	}

	n.Description = description
	n.SetUpdatedAt()
	return nil
}

func (n *Note) SetData(data *NoteData) {
	if data == nil || data.PageURL == "" {
		n.Data = nil
	} else {
		n.Data = data
	}
	n.SetUpdatedAt()
}

func (n *Note) SetUpdatedAt() {
	n.UpdatedAt = manipulation.NowUTC()
}
