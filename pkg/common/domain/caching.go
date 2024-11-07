package domain

import "github.com/namhq1989/go-utilities/appcontext"

type CachingRepository interface {
	GetIpCity(ctx *appcontext.AppContext, ip string) (*string, error)
	SetIpCity(ctx *appcontext.AppContext, ip, city string) error

	GetCityWeather(ctx *appcontext.AppContext, city string) (*Weather, error)
	SetCityWeather(ctx *appcontext.AppContext, city string, weather Weather) error
}
