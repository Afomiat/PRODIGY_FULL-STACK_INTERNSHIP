package usecase

import (
    "context"
    
    "time"

    "github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type AttendanceUsecase struct {
    repo        domain.AttendanceRepository
    userRepo    domain.UserRepository 
}

func NewAttendanceUsecase(repo domain.AttendanceRepository, userRepo domain.UserRepository) domain.AttendanceUsecase {
    return &AttendanceUsecase{repo: repo, userRepo: userRepo}
}

func (uc *AttendanceUsecase) ClockIn(ctx context.Context, userID primitive.ObjectID) error {
    record := &domain.AttendanceRecord{
        ID:       primitive.NewObjectID(),
        UserID:   userID,
        ClockIn:  time.Now(),
    }
    return uc.repo.InsertAttendanceRecord(ctx, record)
}

func (uc *AttendanceUsecase) ClockOut(ctx context.Context, userID primitive.ObjectID) error {
    record, err := uc.repo.FindLatestClockInRecord(ctx, userID)
    if err != nil {
        return err
    }
    clockOutTime := time.Now()
    record.ClockOut = clockOutTime
    return uc.repo.UpdateAttendanceRecord(ctx, record)
}

func (uc *AttendanceUsecase) GetAllAttendanceRecords(ctx context.Context) ([]domain.AttendanceRecord, error) {
    records, err := uc.repo.GetAllRecords(ctx)
    if err != nil {
        return nil, err
    }

    userRecords := make(map[primitive.ObjectID]domain.AttendanceRecord)
    for _, record := range records {
        if existingRecord, exists := userRecords[record.UserID]; !exists || record.ClockIn.After(existingRecord.ClockIn) {
            user, err := uc.userRepo.GetUserByID(ctx, record.UserID)
            if err != nil {
                return nil, err
            }
            record.Username = user.Username
            userRecords[record.UserID] = record
        }
    }

    var result []domain.AttendanceRecord
    for _, record := range userRecords {
        if record.ClockOut.IsZero() {
            record.ClockOut = time.Time{}
        }
        result = append(result, record)
    }

    return result, nil
}
