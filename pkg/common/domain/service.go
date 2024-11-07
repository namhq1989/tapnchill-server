package domain

import "github.com/namhq1989/go-utilities/appcontext"

type Service interface {
	GetCityWeather(ctx *appcontext.AppContext, city string) (*Weather, error)
	GetIpCity(ctx *appcontext.AppContext, ip string) (*string, error)
}
