package domain

import "github.com/namhq1989/go-utilities/appcontext"

type ExternalApiRepository interface {
	GetRandomQuote(ctx *appcontext.AppContext) (*Quote, error)
}
