package infrastructure

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/report"
)

type ReportRepository struct {
	report report.Operations
}

func NewReportRepository(report report.Operations) ReportRepository {
	return ReportRepository{
		report: report,
	}
}

func (r ReportRepository) NewUserSignedInWithGoogle(ctx *appcontext.AppContext, userID, email string) error {
	return r.report.NewUserSignedInWithGoogle(ctx, userID, email)
}
