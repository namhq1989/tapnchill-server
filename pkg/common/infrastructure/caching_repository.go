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
	latestQuoteCachingDuration time.Duration
	ipCityCachingDuration      time.Duration
	cityWeatherCachingDuration time.Duration
}

func NewCachingRepository(caching *caching.Caching) CachingRepository {
	return CachingRepository{
		caching:                    caching,
		domain:                     "common",
		latestQuoteCachingDuration: 3 * time.Hour,
		ipCityCachingDuration:      2 * time.Hour,
		cityWeatherCachingDuration: 2 * time.Hour,
	}
}

// LATEST QUOTE

func (r CachingRepository) GetLatestQuote(ctx *appcontext.AppContext) (*domain.Quote, error) {
	key := r.generateLatestQuoteKey()

	dataStr, err := r.caching.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if dataStr == "" {
		return nil, nil
	}

	var result *domain.Quote
	if err = json.Unmarshal([]byte(dataStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r CachingRepository) SetLatestQuote(ctx *appcontext.AppContext, quote domain.Quote) error {
	key := r.generateLatestQuoteKey()
	r.caching.SetTTL(ctx, key, quote, r.latestQuoteCachingDuration)
	return nil
}

func (r CachingRepository) generateLatestQuoteKey() string {
	return r.caching.GenerateKey(r.domain, "quote:latest")
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
