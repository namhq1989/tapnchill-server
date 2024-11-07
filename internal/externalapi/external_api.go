package externalapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/namhq1989/go-utilities/appcontext"
)

type Operations interface {
	GetRandomQuote(ctx *appcontext.AppContext) (*GetRandomQuoteResult, error)
}

type ExternalApi struct {
	quote *resty.Client
}

const (
	quoteApiEndpoint = "https://quoteslate.vercel.app"
)

func NewExternalAPIClient() *ExternalApi {
	return &ExternalApi{
		quote: resty.New().
			SetBaseURL(quoteApiEndpoint).
			SetHeader("Accept", "application/json").
			SetTimeout(30 * time.Second).
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			SetRetryAfter(func(_ *resty.Client, resp *resty.Response) (time.Duration, error) {
				return 1, fmt.Errorf("failed to send Quote request at %s with status code %d", resp.Request.RawRequest.RequestURI, resp.StatusCode())
			}),
	}
}
