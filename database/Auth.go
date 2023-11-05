package database

import (
	"context"
	"errors"
	"time"

	"github.com/iamtbay/real-estate-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct{}

func InitAuth() *Auth {
	return &Auth{}
}

func (s *Auth) Register(userInfo *models.User) error {
	collection := mongoDB.Client.Database("Go-Real-Estate").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	//check user email exist or not
	filter := bson.M{"email": userInfo.Email}
	result, _ := collection.CountDocuments(ctx, filter)

	documents := models.User{
		Name:     userInfo.Name,
		Surname:  userInfo.Surname,
		Email:    userInfo.Email,
		Password: userInfo.Password,
	}
	//insert user to db
	_, err := collection.InsertOne(ctx, documents)
	if mongo.IsDuplicateKeyError(err) || result > 0 {
		return errors.New("email already in use")
	}

	return nil

}

// Login
func (s *Auth) Login(userInfo *models.User) (*models.User, error) {
	collection := mongoDB.Client.Database("Go-Real-Estate").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	//
	filter := bson.M{"email": userInfo.Email}
	var userFromDB *models.User
	err := collection.FindOne(ctx, filter).Decode(&userFromDB)
	//if email isn't on list
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return nil, errors.New("invalid email or password1 ")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(userInfo.Password))
	//if pass is invalid
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return userFromDB, nil

}

// Logout
func (s *Auth) Logout() {

}
