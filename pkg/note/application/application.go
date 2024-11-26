package application

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/note/application/command"
	"github.com/namhq1989/tapnchill-server/pkg/note/application/query"
	"github.com/namhq1989/tapnchill-server/pkg/note/domain"
	"github.com/namhq1989/tapnchill-server/pkg/note/dto"
)

type (
	Commands interface {
		CreateNote(ctx *appcontext.AppContext, performerID string, req dto.CreateNoteRequest) (*dto.CreateNoteResponse, error)
		UpdateNote(ctx *appcontext.AppContext, performerID, noteID string, req dto.UpdateNoteRequest) (*dto.UpdateNoteResponse, error)
		DeleteNote(ctx *appcontext.AppContext, performerID, noteID string, _ dto.DeleteNoteRequest) (*dto.DeleteNoteResponse, error)
	}
	Queries interface {
		Sync(ctx *appcontext.AppContext, performerID string, req dto.SyncRequest) (*dto.SyncResponse, error)
	}
	Instance interface {
		Commands
		Queries
	}

	commandHandlers struct {
		command.CreateNoteHandler
		command.UpdateNoteHandler
		command.DeleteNoteHandler
	}
	queryHandlers struct {
		query.SyncHandler
	}
	Application struct {
		commandHandlers
		queryHandlers
	}
)

var _ Instance = (*Application)(nil)

func New(
	noteRepository domain.NoteRepository,
	userHub domain.UserHub,
) *Application {
	return &Application{
		commandHandlers: commandHandlers{
			CreateNoteHandler: command.NewCreateNoteHandler(noteRepository, userHub),
			UpdateNoteHandler: command.NewUpdateNoteHandler(noteRepository),
			DeleteNoteHandler: command.NewDeleteNoteHandler(noteRepository),
		},
		queryHandlers: queryHandlers{
			SyncHandler: query.NewSyncHandler(noteRepository),
		},
	}
}
