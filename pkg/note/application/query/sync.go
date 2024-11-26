package query

import (
	"time"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/note/domain"
	"github.com/namhq1989/tapnchill-server/pkg/note/dto"
)

type SyncHandler struct {
	noteRepository domain.NoteRepository
}

func NewSyncHandler(noteRepository domain.NoteRepository) SyncHandler {
	return SyncHandler{
		noteRepository: noteRepository,
	}
}

func (h SyncHandler) Sync(ctx *appcontext.AppContext, performerID string, req dto.SyncRequest) (*dto.SyncResponse, error) {
	ctx.Logger().Info("new sync notes request", appcontext.Fields{
		"performerID": performerID, "lastUpdatedAt": req.LastUpdatedAt,
	})

	ctx.Logger().Text("parse time")
	var (
		lastUpdatedAt = time.Time{}
		err           error
	)
	if req.LastUpdatedAt != "" {
		lastUpdatedAt, err = time.Parse(time.RFC3339, req.LastUpdatedAt)
		if err != nil {
			ctx.Logger().Error("failed to parse time", err, appcontext.Fields{})
			return nil, err
		}
	}

	ctx.Logger().Text("find in db")
	notes, err := h.noteRepository.Sync(ctx, performerID, lastUpdatedAt, domain.NumOfNotesEachSync)
	if err != nil {
		ctx.Logger().Error("failed to find in db", err, appcontext.Fields{})
		return nil, err
	}

	totalNotes := len(notes)
	if totalNotes == 0 {
		ctx.Logger().Text("done sync notes request")
		return &dto.SyncResponse{
			Notes: make([]dto.Note, 0),
			Limit: domain.NumOfNotesEachSync,
		}, nil
	}

	ctx.Logger().Text("convert response data")
	result := make([]dto.Note, 0)
	for _, note := range notes {
		result = append(result, dto.Note{}.FromDomain(note))
	}

	ctx.Logger().Text("done sync notes request")
	return &dto.SyncResponse{
		Notes: result,
		Limit: domain.NumOfNotesEachSync,
	}, nil
}
