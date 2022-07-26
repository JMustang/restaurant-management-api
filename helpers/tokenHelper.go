package helpers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/krishpranav/golang-management/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type SingnedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
