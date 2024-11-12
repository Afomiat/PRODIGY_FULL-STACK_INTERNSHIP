package employeeUtil

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/Afomiat/PRODIGY_FULL-STACK_INTERNSHIP/domain"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

func ValidateEmail(email string) bool {

	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}

func ValidatePassword(password string) bool {
	return len(password) >= 8
}

func GenerateOTP() string {
	rand.Seed(uint64(time.Now().UnixNano())) // Seed the random number generator with int64
	otp := rand.Intn(900000) + 100000        // Generate a random number between 100000 and 999999
	return fmt.Sprintf("%06d", otp)          // Format the number as a 6-digit string
}
func HassPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) error {
    fmt.Println("Hashed Password:", hashedPassword)
    fmt.Println("Input Password:", password)

    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        fmt.Println("Error comparing passwords:", err)
        return fmt.Errorf("invalid password")
    }
    return nil
}

func CanManipulateUser(claims *domain.JwtCustomClaims, user *domain.User, manip string) *domain.Error {
	// If the user is a regular user, they can only manipulate their own account.
	if claims.Role == "user" {
		if user.ID != claims.UserID {
			var message string
			if manip == "add" {
				message = "A User cannot add a new user"
			} else {
				message = "A User cannot " + manip + " another user"
			}

			return &domain.Error{
				Err:        errors.New("unauthorized"),
				StatusCode: http.StatusForbidden,
				Message:    message,
			}
		}

		return nil
	}

	// If the user is an admin, they can manipulate all users except root user and other admin users.
	if claims.Role == "admin" {
		if user.Role == "root" {
			return &domain.Error{
				Err:        errors.New("forbidden"),
				StatusCode: http.StatusForbidden,
				Message:    "Cannot " + manip + " root user",
			}
		}

		if user.Role == "admin" && claims.UserID != user.ID {
			return &domain.Error{
				Err:        errors.New("unauthorized"),
				StatusCode: http.StatusForbidden,
				Message:    "Admin cannot " + manip + " another admin user",
			}
		}
	}

	// If the user is a root user, they can manipulate all users.
	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
