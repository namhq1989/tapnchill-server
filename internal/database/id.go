package database

import "go.mongodb.org/mongo-driver/bson/primitive"

func NewStringID() string {
	return primitive.NewObjectID().Hex()
}

func NewObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}

func ObjectIDFromString(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

func IsValidObjectID(id string) bool {
	_, err := ObjectIDFromString(id)
	return err == nil
}
