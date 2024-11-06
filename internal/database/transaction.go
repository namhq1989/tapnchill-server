package database

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func (db Database) Transaction(ctx *appcontext.AppContext, fn func(ssCtx mongo.SessionContext) (interface{}, error)) error {
	tnxOpts := options.Transaction().SetWriteConcern(writeconcern.Majority())
	sess, err := db.mongo.Client().StartSession()
	if err != nil {
		return err
	}
	defer sess.EndSession(ctx.Context())

	_, err = sess.WithTransaction(ctx.Context(), fn, tnxOpts)
	return err
}
