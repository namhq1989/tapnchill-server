package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/tapnchill-server/internal/error"
	"github.com/namhq1989/tapnchill-server/pkg/note/domain"
	"github.com/namhq1989/tapnchill-server/pkg/note/dto"
)

type CreateNoteHandler struct {
	noteRepository domain.NoteRepository
	userHub        domain.UserHub
}

func NewCreateNoteHandler(noteRepository domain.NoteRepository, userHub domain.UserHub) CreateNoteHandler {
	return CreateNoteHandler{
		noteRepository: noteRepository,
		userHub:        userHub,
	}
}

func (h CreateNoteHandler) CreateNote(ctx *appcontext.AppContext, performerID string, req dto.CreateNoteRequest) (*dto.CreateNoteResponse, error) {
	ctx.Logger().Info("new create note request", appcontext.Fields{
		"performerID": performerID, "title": req.Title, "description": req.Description, "data": req.Data,
	})

	ctx.Logger().Text("get user note quota")
	quota, err := h.userHub.GetNoteQuota(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to get user note quota", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("count user total notes")
	totalNotes, err := h.noteRepository.CountByUserID(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to count user total notes", err, appcontext.Fields{})
		return nil, err
	}

	if totalNotes >= quota {
		ctx.Logger().Error("user note quota exceeded", err, appcontext.Fields{"quota": quota, "total": totalNotes})
		return nil, apperrors.User.ResourceLimitReached
	}

	var data *domain.NoteData
	if req.Data != nil {
		data = &domain.NoteData{
			PageText:  req.Data.PageText,
			PageTitle: req.Data.PageTitle,
			PageURL:   req.Data.PageURL,
		}
	}

	ctx.Logger().Text("create new note model")
	note, err := domain.NewNote(performerID, req.Title, req.Description, data)
	if err != nil {
		ctx.Logger().Error("failed to create new note model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist note in db")
	if err = h.noteRepository.Create(ctx, *note); err != nil {
		ctx.Logger().Error("failed to persist note in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done create note request")
	return &dto.CreateNoteResponse{
		ID: note.ID,
	}, nil
}
