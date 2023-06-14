package db

import (
	"context"
	"github.com/jafari-mohammad-reza/hotel-reservation.git/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	MongoDbAbstractRepository[types.User]
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		MongoDbAbstractRepository: MongoDbAbstractRepository[types.User]{Collection: collection},
	}
}

func (repo *UserRepository) Login(ctx context.Context, email string, password string) (*primitive.ObjectID, error) {
	result := repo.Collection.FindOne(ctx, bson.M{"email": email})
	var user types.User
	decodeErr := result.Decode(&user)
	if decodeErr != nil {
		return nil, decodeErr
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return &user.ID, nil
}
