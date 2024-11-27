package usecase

import (
    "context"
    "time"

    "github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type AttendanceUsecase struct {
    repo domain.AttendanceRepository
}

func NewAttendanceUsecase(repo domain.AttendanceRepository) domain.AttendanceUsecase {
    return &AttendanceUsecase{repo: repo}
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
    record.ClockOut = time.Now()
    return uc.repo.UpdateAttendanceRecord(ctx, record)
}

func (uc *AttendanceUsecase) GetAllAttendanceRecords(ctx context.Context) ([]domain.AttendanceRecord, error) {
    return uc.repo.GetAllRecords(ctx)
}
