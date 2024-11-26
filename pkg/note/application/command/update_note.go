package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/note/domain"
	"github.com/namhq1989/tapnchill-server/pkg/note/dto"
)

type UpdateNoteHandler struct {
	noteRepository domain.NoteRepository
}

func NewUpdateNoteHandler(noteRepository domain.NoteRepository) UpdateNoteHandler {
	return UpdateNoteHandler{
		noteRepository: noteRepository,
	}
}

func (h UpdateNoteHandler) UpdateNote(ctx *appcontext.AppContext, performerID, noteID string, req dto.UpdateNoteRequest) (*dto.UpdateNoteResponse, error) {
	ctx.Logger().Info("new update note request", appcontext.Fields{
		"performerID": performerID, "noteID": noteID,
		"title": req.Title, "description": req.Description, "data": req.Data,
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

	ctx.Logger().Text("set note data")
	if err = note.SetTitle(req.Title); err != nil {
		ctx.Logger().Error("failed to set note title", err, appcontext.Fields{})
		return nil, err
	}
	if err = note.SetDescription(req.Description); err != nil {
		ctx.Logger().Error("failed to set note description", err, appcontext.Fields{})
		return nil, err
	}
	if req.Data != nil {
		note.SetData(&domain.NoteData{
			PageText:  req.Data.PageText,
			PageTitle: req.Data.PageTitle,
			PageURL:   req.Data.PageURL,
		})
	} else {
		note.SetData(nil)
	}

	ctx.Logger().Text("update note in db")
	if err = h.noteRepository.Update(ctx, *note); err != nil {
		ctx.Logger().Error("failed to update note in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done update note request")
	return &dto.UpdateNoteResponse{
		ID: note.ID,
	}, nil
}
