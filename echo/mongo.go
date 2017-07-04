package echo

import (
	mgo "gopkg.in/mgo.v2"

	"github.com/labstack/echo"
)

const (
	// MongoSessionKey is the key used to set the MongoDB session value
	// in the Echo request, us it to retreive that value.
	MongoSessionKey = "mongo-session-key"
)

// MongoMiddleware takes a MongoDB session and ensures that it's copied
// to subsequent handlers.
func MongoMiddleware(session *mgo.Session) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			copy := session.Copy()
			defer copy.Close()

			c.Set(MongoSessionKey, copy)

			return next(c)
		}
	}
}
