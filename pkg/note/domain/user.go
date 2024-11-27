package domain

import (
	"github.com/namhq1989/go-utilities/appcontext"
)

type UserHub interface {
	GetNoteQuota(ctx *appcontext.AppContext, userID string) (int64, error)
}