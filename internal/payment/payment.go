package payment

import (
	paddle "github.com/PaddleHQ/paddle-go-sdk/v2"
)

type Operations interface {
}

type Payment struct {
	paddle *paddle.SDK
}

func NewPayment(paddleApiKey string, isEnvRelease bool) *Payment {
	paddleBaseURL := paddle.SandboxBaseURL
	if isEnvRelease {
		paddleBaseURL = paddle.ProductionBaseURL
	}

	paddleSdk, err := paddle.New(
		paddleApiKey,
		paddle.WithBaseURL(paddleBaseURL),
	)
	if err != nil {
		panic(err)
	}

	return &Payment{
		paddle: paddleSdk,
	}
}
