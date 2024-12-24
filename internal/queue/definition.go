package queue

var TypeNames = struct {
	CreateUserDefaultGoal string

	GetRandomQuote string

	DowngradeExpiredSubscriptions string

	PaddleSubscriptionCreated              string
	PaddleTransactionCompleted             string
	LemonsqueezySubscriptionPaymentSuccess string
}{
	CreateUserDefaultGoal: "task.createUserDefaultGoal",

	GetRandomQuote: "common.getRandomQuote",

	DowngradeExpiredSubscriptions: "user.downgradeExpiredSubscriptions",

	PaddleSubscriptionCreated:              "webhook.paddleSubscriptionCreated",
	PaddleTransactionCompleted:             "webhook.paddleTransactionCompleted",
	LemonsqueezySubscriptionPaymentSuccess: "webhook.lemonsqueezySubscriptionPaymentSuccess",
}
