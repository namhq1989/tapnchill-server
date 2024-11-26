package dbmodel

type NoteData struct {
	PageText  string `bson:"pageText"`
	PageTitle string `bson:"pageTitle"`
	PageURL   string `bson:"pageUrl"`
}
