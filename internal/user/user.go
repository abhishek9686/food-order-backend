package user

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User ...
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string    `gorm:"size:255;not null;unique" json:"first_name"`
	LastName  string    `gorm:"size:255;not null;unique" json:"last_name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Phone     int64     `gorm:"size:100" json:"phone"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Address   string    `gorm:"size:255" json:"address"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Hash - to encrypt user's password.
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword - compare hashed password and given password.
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// BeforeSave ...
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Prepare ...
func (u *User) Prepare() {
	u.ID = 0
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Address = html.EscapeString(strings.TrimSpace(u.Address))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// Validate ...
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.FirstName == "" {
			return errors.New("Required Firstname")
		}
		if u.LastName == "" {
			return errors.New("Required Lastname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.FirstName == "" {
			return errors.New("Required Firstname")
		}
		if u.LastName == "" {
			return errors.New("Required Lastname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

// CreateUser - add new user to DB table.
func (u *User) CreateUser(db *gorm.DB) (*User, error) {
	var err error
	u.Prepare()
	u.BeforeSave()
	if err := u.Validate(""); err != nil {
		return &User{}, err
	}
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
