package queue

var TypeNames = struct {
	CreateUserDefaultGoal string

	GetRandomQuote string

	PaddleSubscriptionCreated       string
	PaddleTransactionCompleted      string
	FastspringSubscriptionActivated string
}{
	CreateUserDefaultGoal: "task.createUserDefaultGoal",

	GetRandomQuote: "common.getRandomQuote",

	PaddleSubscriptionCreated:       "webhook.paddleSubscriptionCreated",
	PaddleTransactionCompleted:      "webhook.paddleTransactionCompleted",
	FastspringSubscriptionActivated: "webhook.fastspringSubscriptionActivated",
}
