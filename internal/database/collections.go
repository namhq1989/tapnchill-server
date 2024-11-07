package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	User string

	Quote    string
	Feedback string

	Goal string
	Task string
}{
	User: "user.users",

	Quote:    "common.quotes",
	Feedback: "common.feedbacks",

	Goal: "task.goals",
	Task: "task.tasks",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
