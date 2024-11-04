package repository

import (
	"context"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginRepo struct {
	collection *mongo.Collection
}

func NewLoginRepo(db *mongo.Database, coll string) *LoginRepo {
	return &LoginRepo{
		collection: db.Collection(coll),
	}
}

func (u *LoginRepo) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	filter := bson.M{"email": email}
	user := &domain.User{}
	err := u.collection.FindOne(c, filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
