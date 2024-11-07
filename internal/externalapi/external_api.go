package externalapi

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/namhq1989/go-utilities/appcontext"
)

type Operations interface {
	GetRandomQuote(ctx *appcontext.AppContext) (*GetRandomQuoteResult, error)
	GetIpCity(ctx *appcontext.AppContext, ip string) (*GetIpCityResult, error)
	GetCityWeather(ctx *appcontext.AppContext, city string) (*GetCityWeatherResult, error)
}

type ExternalApi struct {
	visualCrossingToken string
	ipInfoToken         string

	quoteClient    *resty.Client
	weatherClient  *resty.Client
	locationClient *resty.Client
}

const (
	quoteApiEndpoint    = "https://quoteslate.vercel.app"
	weatherApiEndpoint  = "https://weather.visualcrossing.com"
	locationApiEndpoint = "https://ipinfo.io"
)

func NewExternalAPIClient(visualCrossingToken string, ipInfoToken string) *ExternalApi {
	return &ExternalApi{
		visualCrossingToken: visualCrossingToken,
		ipInfoToken:         ipInfoToken,

		quoteClient: resty.New().
			SetBaseURL(quoteApiEndpoint).
			SetHeader("Accept", "application/json").
			SetTimeout(30 * time.Second).
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			SetRetryAfter(func(_ *resty.Client, resp *resty.Response) (time.Duration, error) {
				return 1, fmt.Errorf("failed to send Quote request at %s with status code %d", resp.Request.RawRequest.RequestURI, resp.StatusCode())
			}),
		weatherClient: resty.New().
			SetBaseURL(weatherApiEndpoint).
			SetHeader("Accept", "application/json").
			SetTimeout(30 * time.Second).
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			SetRetryAfter(func(_ *resty.Client, resp *resty.Response) (time.Duration, error) {
				return 1, fmt.Errorf("failed to send Weather request at %s with status code %d", resp.Request.RawRequest.RequestURI, resp.StatusCode())
			}),
		locationClient: resty.New().
			SetBaseURL(locationApiEndpoint).
			SetHeader("Accept", "application/json").
			SetTimeout(30 * time.Second).
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			SetRetryAfter(func(_ *resty.Client, resp *resty.Response) (time.Duration, error) {
				return 1, fmt.Errorf("failed to send Location request at %s with status code %d", resp.Request.RawRequest.RequestURI, resp.StatusCode())
			}),
	}
}
