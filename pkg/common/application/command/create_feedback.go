package command

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
	"github.com/namhq1989/tapnchill-server/pkg/common/dto"
)

type CreateFeedbackHandler struct {
	feedbackRepository domain.FeedbackRepository
	reportRepository   domain.ReportRepository
}

func NewCreateFeedbackHandler(feedbackRepository domain.FeedbackRepository, reportRepository domain.ReportRepository) CreateFeedbackHandler {
	return CreateFeedbackHandler{
		feedbackRepository: feedbackRepository,
		reportRepository:   reportRepository,
	}
}

func (h CreateFeedbackHandler) CreateFeedback(ctx *appcontext.AppContext, performerID string, req dto.CreateFeedbackRequest) (*dto.CreateFeedbackResponse, error) {
	ctx.Logger().Info("new create feedback request", appcontext.Fields{"performerID": performerID, "email": req.Email, "feedback": req.Feedback})

	ctx.Logger().Text("create new feedback model")
	feedback, err := domain.NewFeedback(performerID, req.Email, req.Feedback, ctx.GetIP())
	if err != nil {
		ctx.Logger().Error("failed to create new feedback model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("persist feedback in db")
	if err = h.feedbackRepository.Create(ctx, *feedback); err != nil {
		ctx.Logger().Error("failed to persist feedback in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("send report for new user signed in with Google")
	go func() {
		if err = h.reportRepository.NewUserFeedback(ctx, feedback.ID, feedback.Feedback); err != nil {
			ctx.Logger().Error("failed to send report", err, appcontext.Fields{})
		}
	}()

	ctx.Logger().Text("done create feedback request")
	return &dto.CreateFeedbackResponse{
		ID: feedback.ID,
	}, nil
}
