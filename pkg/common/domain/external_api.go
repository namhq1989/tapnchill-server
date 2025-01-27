package domain

import "github.com/namhq1989/go-utilities/appcontext"

type ExternalApiRepository interface {
	GetRandomQuote(ctx *appcontext.AppContext) (*Quote, error)
	GetIpCity(ctx *appcontext.AppContext, ip string) (*string, error)
	GetCityWeather(ctx *appcontext.AppContext, city string) (*Weather, error)
}
