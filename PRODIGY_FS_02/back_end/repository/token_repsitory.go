package repository

import (
	"context"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTokenRepository struct {
	collection         *mongo.Collection
	
}

func NewTokenRepository(db *mongo.Database, coll string) domain.TokenRepository {
	return &MongoTokenRepository{
		collection: db.Collection(coll),
	}
}

func (repo *MongoTokenRepository) SaveToken(ctx context.Context, token *domain.Token) error {
	_, err := repo.collection.InsertOne(ctx,token)
	return err
}

func (repo *MongoTokenRepository) FindTokenByRefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error) {
	var token domain.Token
	err := repo.collection.FindOne(ctx, bson.M{"refresh_token": refreshToken}).Decode(&token)
	return &token, err
}