package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/note/domain"
	"github.com/namhq1989/tapnchill-server/pkg/note/dto"
)

type DeleteNoteHandler struct {
	noteRepository domain.NoteRepository
}

func NewDeleteNoteHandler(noteRepository domain.NoteRepository) DeleteNoteHandler {
	return DeleteNoteHandler{
		noteRepository: noteRepository,
	}
}

func (h DeleteNoteHandler) DeleteNote(ctx *appcontext.AppContext, performerID, noteID string, _ dto.DeleteNoteRequest) (*dto.DeleteNoteResponse, error) {
	ctx.Logger().Info("new delete note request", appcontext.Fields{
		"performerID": performerID, "noteID": noteID,
	})

	note, err := h.noteRepository.FindByID(ctx, noteID)
	if err != nil {
		ctx.Logger().Error("failed to find note in db", err, appcontext.Fields{})
		return nil, err
	}
	if note == nil {
		ctx.Logger().ErrorText("note not found, respond")
		return nil, apperrors.Common.NotFound
	}
	if note.UserID != performerID {
		ctx.Logger().ErrorText("note author not match, respond")
		return nil, apperrors.Common.NotFound
	}

	ctx.Logger().Text("delete note in db")
	if err = h.noteRepository.Delete(ctx, note.ID); err != nil {
		ctx.Logger().Error("failed to delete note in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done delete note request")
	return &dto.DeleteNoteResponse{}, nil
}
