package config

import (
	"errors"
)

type (
	Server struct {
		RestPort string
		GRPCPort string

		AppName      string
		Environment  string
		IsEnvRelease bool

		// Authentication
		FirebaseServiceAccount      string
		AnonymousUserChecksumSecret string
		AccessTokenSecret           string
		AccessTokenTTL              int

		// Mongo
		MongoURL    string
		MongoDBName string

		// Redis
		CachingRedisURL string

		// Queue
		QueueRedisURL    string
		QueueUsername    string
		QueuePassword    string
		QueueConcurrency int

		// Lemonsqueezy
		LemonsqueezyAPIToken                        string
		LemonsqueezySigningSecret                   string
		LemonsqueezyStoreID                         string
		LemonsqueezySubscriptionMonthlyVariantID    string
		LemonsqueezySubscriptionMonthlyDiscountCode string
		LemonsqueezySubscriptionYearlyVariantID     string
		LemonsqueezySubscriptionYearlyDiscountCode  string

		// 3rd party
		IpInfoToken         string
		VisualCrossingToken string
		TelegramBotToken    string
		TelegramChannelID   string
	}
)

func Init() Server {
	cfg := Server{
		RestPort: ":3070",
		GRPCPort: ":3071",

		AppName:     getEnvStr("APP_NAME"),
		Environment: getEnvStr("ENVIRONMENT"),

		FirebaseServiceAccount:      getEnvStr("FIREBASE_SERVICE_ACCOUNT"),
		AnonymousUserChecksumSecret: getEnvStr("ANONYMOUS_USER_CHECKSUM_SECRET"),
		AccessTokenSecret:           getEnvStr("ACCESS_TOKEN_SECRET"),
		AccessTokenTTL:              getEnvInt("ACCESS_TOKEN_TTL"),

		MongoURL:    getEnvStr("MONGO_URL"),
		MongoDBName: getEnvStr("MONGO_DB_NAME"),

		CachingRedisURL: getEnvStr("CACHING_REDIS_URL"),

		QueueRedisURL:    getEnvStr("QUEUE_REDIS_URL"),
		QueueUsername:    getEnvStr("QUEUE_USERNAME"),
		QueuePassword:    getEnvStr("QUEUE_PASSWORD"),
		QueueConcurrency: getEnvInt("QUEUE_CONCURRENCY"),

		LemonsqueezyAPIToken:                        getEnvStr("LEMONSQUEEZY_API_TOKEN"),
		LemonsqueezySigningSecret:                   getEnvStr("LEMONSQUEEZY_SIGNING_SECRET"),
		LemonsqueezyStoreID:                         getEnvStr("LEMONSQUEEZY_STORE_ID"),
		LemonsqueezySubscriptionMonthlyVariantID:    getEnvStr("LEMONSQUEEZY_SUBSCRIPTION_MONTHLY_VARIANT_ID"),
		LemonsqueezySubscriptionMonthlyDiscountCode: getEnvStr("LEMONSQUEEZY_SUBSCRIPTION_MONTHLY_DISCOUNT_CODE"),
		LemonsqueezySubscriptionYearlyVariantID:     getEnvStr("LEMONSQUEEZY_SUBSCRIPTION_YEARLY_VARIANT_ID"),
		LemonsqueezySubscriptionYearlyDiscountCode:  getEnvStr("LEMONSQUEEZY_SUBSCRIPTION_YEARLY_DISCOUNT_CODE"),

		IpInfoToken:         getEnvStr("IP_INFO_TOKEN"),
		VisualCrossingToken: getEnvStr("VISUAL_CROSSING_TOKEN"),
		TelegramBotToken:    getEnvStr("TELEGRAM_BOT_TOKEN"),
		TelegramChannelID:   getEnvStr("TELEGRAM_CHANNEL_ID"),
	}
	cfg.IsEnvRelease = cfg.Environment == "release"

	// validation
	if cfg.Environment == "" {
		panic(errors.New("missing ENVIRONMENT"))
	}

	if cfg.FirebaseServiceAccount == "" {
		panic(errors.New("missing FIREBASE_SERVICE_ACCOUNT"))
	}

	if cfg.AccessTokenSecret == "" {
		panic(errors.New("missing ACCESS_TOKEN_SECRET"))
	}

	if cfg.MongoURL == "" {
		panic(errors.New("missing MONGO_URL"))
	}

	if cfg.CachingRedisURL == "" {
		panic(errors.New("missing CACHING_REDIS_URL"))
	}

	if cfg.QueueRedisURL == "" {
		panic(errors.New("missing QUEUE_REDIS_URL"))
	}

	if cfg.LemonsqueezyAPIToken == "" {
		panic(errors.New("missing LEMONSQUEEZY_API_TOKEN"))
	}

	if cfg.LemonsqueezySigningSecret == "" {
		panic(errors.New("missing LEMONSQUEEZY_SIGNING_SECRET"))
	}

	if cfg.LemonsqueezyStoreID == "" {
		panic(errors.New("missing LEMONSQUEEZY_STORE_ID"))
	}

	return cfg
}
