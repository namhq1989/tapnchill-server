package report

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/namhq1989/go-utilities/appcontext"
	"time"
)

type Operations interface {
	NewUserSignedInWithGoogle(ctx *appcontext.AppContext, userID, email string) error
	NewUserFeedback(ctx *appcontext.AppContext, feedbackID, content string) error
}

type Report struct {
	telegramBotToken string
	telegramChatID   string
	telegramClient   *resty.Client
}

const (
	telegramApiEndpoint = "https://api.telegram.org"
)

func NewReport(telegramBotToken, telegramChatID string) *Report {
	return &Report{
		telegramBotToken: telegramBotToken,
		telegramChatID:   telegramChatID,
		telegramClient: resty.New().
			SetBaseURL(telegramApiEndpoint).
			SetHeader("Accept", "application/json").
			SetTimeout(30 * time.Second).
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			SetRetryAfter(func(_ *resty.Client, resp *resty.Response) (time.Duration, error) {
				return 1, fmt.Errorf("failed to send Telegram request at %s with status code %d", resp.Request.RawRequest.RequestURI, resp.StatusCode())
			}),
	}
}
