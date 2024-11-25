package repository

import (
	"context"
	"fmt"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	collection *mongo.Collection
}

func (u *EmployeeRepo) GetUserByID(c context.Context, id primitive.ObjectID) (*domain.User, error) {
	filter := bson.M{"_id": id}
	user := &domain.User{}
	err := u.collection.FindOne(c, filter).Decode(user)
	return user, err
}

func (u *EmployeeRepo) UpdateUser(c context.Context, user *domain.User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	_, err := u.collection.UpdateOne(c, filter, update)
	return err
}



func NewEmployeeRepo(db *mongo.Database, coll string) *EmployeeRepo {
	return &EmployeeRepo{
		collection: db.Collection(coll),
	}
}

func (u *EmployeeRepo) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	filter := bson.M{"email": email}
	user := &domain.User{}
	err := u.collection.FindOne(c, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	fmt.Print(user, "user in get user by email /////")
	return user, nil
}

func (u *EmployeeRepo) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {

	var user domain.User

	filter := bson.D{{Key: "username", Value: username}}
	err := u.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil 
		}
		return nil, err 
	}
	return &user, nil
}

func (u *EmployeeRepo) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := u.collection.InsertOne(ctx, user)

	return err
}

func (u *EmployeeRepo) DeleteUser(c context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := u.collection.DeleteOne(c, filter)
	return err
}

func (u *EmployeeRepo) GetAllUsers(c context.Context) ([]*domain.User, error) {
	cursor, err := u.collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)
	
	var users []*domain.User
	for cursor.Next(c) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

