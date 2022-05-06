package user

import (
	"database/sql"
	"errors"
	"fmt"
	"html"
	"strings"
	"sync"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ActiveUserSessionMap = make(map[LoginReq]time.Time)
	UserSessionMapMutex  = &sync.RWMutex{}
)

// LoginReq ...
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User ...
type User struct {
	Role      string    `gorm:"default:user" json:"role"`
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string    `gorm:"size:255;not null" json:"first_name"`
	LastName  string    `gorm:"size:255;not null" json:"last_name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Phone     int64     `gorm:"size:100;unique;not null" json:"phone"`
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
func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		tx.Model(u).Update("role", "admin")
	}
	return
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
	if err := u.Validate(""); err != nil {
		return &User{}, err
	}
	tx := db.Debug().Create(&u)
	if tx != nil && tx.Error != nil {
		err = tx.Error
	}
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// CheckIfValidUser ...
func CheckIfValidUser(db *gorm.DB, loginDetails LoginReq) error {
	var user User
	UserSessionMapMutex.RLock()
	if _, ok := ActiveUserSessionMap[loginDetails]; ok {
		UserSessionMapMutex.RUnlock()
		return nil
	}
	UserSessionMapMutex.RUnlock()

	tx := db.Debug().Raw("SELECT role,id FROM user WHERE email = @email ",
		sql.Named("email", loginDetails.Email)).Find(&user)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("user with email '%s' not found", loginDetails.Email)
	}
	if tx.Error != nil {
		return tx.Error
	}
	err := VerifyPassword(user.Password, loginDetails.Password)
	if err != nil {
		return errors.New("password entered is incorrect")
	}
	UserSessionMapMutex.Lock()
	ActiveUserSessionMap[loginDetails] = time.Now()
	UserSessionMapMutex.Unlock()
	return nil
}
