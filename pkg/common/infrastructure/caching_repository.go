package infrastructure

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/tapnchill-server/internal/caching"
	"github.com/namhq1989/tapnchill-server/pkg/common/domain"
)

type CachingRepository struct {
	caching caching.Operations

	domain                     string
	ipCityCachingDuration      time.Duration
	cityWeatherCachingDuration time.Duration
}

func NewCachingRepository(caching *caching.Caching) CachingRepository {
	return CachingRepository{
		caching:                    caching,
		domain:                     "common",
		ipCityCachingDuration:      1 * time.Hour,
		cityWeatherCachingDuration: 1 * time.Hour,
	}
}

// IP CITY

func (r CachingRepository) GetIpCity(ctx *appcontext.AppContext, ip string) (*string, error) {
	key := r.generateIpCityKey(ip)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if dataStr == "" {
		return nil, nil
	}

	var result string
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r CachingRepository) SetIpCity(ctx *appcontext.AppContext, ip, city string) error {
	key := r.generateIpCityKey(ip)
	r.caching.SetTTL(ctx, key, city, r.ipCityCachingDuration)
	return nil
}

func (r CachingRepository) generateIpCityKey(ip string) string {
	return r.caching.GenerateKey(r.domain, fmt.Sprintf("ip:%s:city", ip))
}

// CITY WEATHER

func (r CachingRepository) GetCityWeather(ctx *appcontext.AppContext, city string) (*domain.Weather, error) {
	key := r.generateCityWeatherKey(city)

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if dataStr == "" {
		return nil, nil
	}

	var result *domain.Weather
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, nil
	}
	return result, nil
}

func (r CachingRepository) SetCityWeather(ctx *appcontext.AppContext, city string, weather domain.Weather) error {
	key := r.generateCityWeatherKey(city)
	r.caching.SetTTL(ctx, key, weather, r.cityWeatherCachingDuration)
	return nil
}

func (r CachingRepository) generateCityWeatherKey(city string) string {
	return r.caching.GenerateKey(r.domain, fmt.Sprintf("city:%s:weather", city))
}
