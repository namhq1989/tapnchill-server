package queue

var TypeNames = struct {
	CreateUserDefaultGoal string

	GetRandomQuote string

	SubscriptionCreated  string
	TransactionCompleted string
}{
	CreateUserDefaultGoal: "task.createUserDefaultGoal",

	GetRandomQuote: "common.getRandomQuote",

	SubscriptionCreated:  "webhook.subscriptionCreated",
	TransactionCompleted: "webhook.transactionCompleted",
}
