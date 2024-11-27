package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	Attendance = "attendance_records"
	CollectionUser      = "employees"
	AdminRole      Role = "ADMIN"
	EmployeeRole       Role = "EMPLOYEE"
)
type ContactInfo struct {
	Address      string `bson:"address"`
	Phone_number string `bson:"phone_number" json:"phone_number"`
}

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	First_Name      string             `bson:"first_name" json:"first_name"`
	Last_Name       string             `bson:"last_name" json:"last_name" `
	Username        string             `bson:"username" json:"username"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password" `
	Role            Role               `bson:"role" json:"role"`
	Bio             string             `bson:"bio" json:"bio"`
	Profile_Picture string             `bson:"profile_picture" json:"profile_picture"`
	Contact_Info    []ContactInfo      `bson:"contact_info" json:"contact_info"`
}

type Privilage struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     Role   `json:"role"`
}




type AttendanceRecord struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    UserID   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Username        string             `bson:"username" json:"username"`
    ClockIn  time.Time          `bson:"clock_in" json:"clock_in"`
    ClockOut time.Time          `bson:"clock_out,omitempty" json:"clock_out,omitempty"`
}
