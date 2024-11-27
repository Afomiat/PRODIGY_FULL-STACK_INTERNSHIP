package repository

import (
    "context"
    
    "github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type AttendanceRepository struct {
    collection *mongo.Collection
}

func NewAttendanceRepository(db *mongo.Database, collectionName string) domain.AttendanceRepository {
    return &AttendanceRepository{
        collection: db.Collection(collectionName),
    }
}

func (repo *AttendanceRepository) InsertAttendanceRecord(ctx context.Context, record *domain.AttendanceRecord) error {
    _, err := repo.collection.InsertOne(ctx, record)
    return err
}

func (repo *AttendanceRepository) FindLatestClockInRecord(ctx context.Context, userID primitive.ObjectID) (*domain.AttendanceRecord, error) {
    var record domain.AttendanceRecord
    filter := bson.M{"user_id": userID, "clock_out": nil}
    err := repo.collection.FindOne(ctx, filter).Decode(&record)
    return &record, err
}

func (repo *AttendanceRepository) UpdateAttendanceRecord(ctx context.Context, record *domain.AttendanceRecord) error {
    filter := bson.M{"_id": record.ID}
    update := bson.M{"$set": record}
    _, err := repo.collection.UpdateOne(ctx, filter, update)
    return err
}

func (repo *AttendanceRepository) GetAllRecords(ctx context.Context) ([]domain.AttendanceRecord, error) {
    var records []domain.AttendanceRecord
    cursor, err := repo.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var record domain.AttendanceRecord
        if err := cursor.Decode(&record); err != nil {
            return nil, err
        }
        records = append(records, record)
    }
    return records, nil
}
