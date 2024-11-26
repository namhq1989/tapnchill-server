package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	User                string
	SubscriptionHistory string

	Quote    string
	Feedback string

	Goal string
	Task string

	Habit           string
	HabitCompletion string
	HabitDailyStats string

	Note string
}{
	User:                "user.users",
	SubscriptionHistory: "user.subscriptionHistories",

	Quote:    "common.quotes",
	Feedback: "common.feedbacks",

	Goal: "task.goals",
	Task: "task.tasks",

	Habit:           "habit.habits",
	HabitCompletion: "habit.completions",
	HabitDailyStats: "habit.dailyStats",

	Note: "note.notes",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
