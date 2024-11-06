package database

import "go.mongodb.org/mongo-driver/mongo"

var Collections = struct {
	User string

	Quote    string
	Feedback string
}{
	User: "user.users",

	Quote:    "common.quotes",
	Feedback: "common.feedbacks",
}

func (db Database) GetCollection(table string) *mongo.Collection {
	return db.mongo.Collection(table)
}
