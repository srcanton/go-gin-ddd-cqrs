package user_aggregate

import (
	"fmt"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/email"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/first_name"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/id"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/last_name"
	"github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/password"
	"time"
)

// User is an user which has id, email and person info
type User struct {
	ID        user_id.UserId            `gorm:"primary_key;column:id;type:uuid" json:"id"`
	FirstName user_first_name.FirstName `gorm:"column:first_name;type:text;size:100;not null;" json:"first_name"`
	LastName  user_last_name.LastName   `gorm:"column:last_name;type:text;size:100;not null;" json:"last_name"`
	Email     user_email.Email          `gorm:"column:email;type:text;size:100;not null;" json:"email"`
	Password  user_password.Password    `gorm:"column:password;type:text;size:100;not null;" json:"password"`
	CreatedAt time.Time                 `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time                 `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time                `json:"deleted_at,omitempty"`
}

func FromRaw(array map[string]interface{}) (*User, error) {
	userIdValueObject, _ := user_id.FromValue(fmt.Sprintf("%v", array["ID"]))
	userFirstNameValueObject, _ := user_first_name.FromValue(fmt.Sprintf("%v", array["FirstName"]))
	userLastNameValueObject, _ := user_last_name.FromValue(fmt.Sprintf("%v", array["LastName"]))
	userEmailValueObject, _ := user_email.FromValue(fmt.Sprintf("%v", array["Email"]))
	userPasswordValueObject, _ := user_password.FromValue(fmt.Sprintf("%v", array["Password"]))
	CreatedAt, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%v", array["CreatedAt"]))
	UpdatedAt, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%v", array["UpdatedAt"]))
	return &User{
		ID:        userIdValueObject,
		FirstName: userFirstNameValueObject,
		LastName:  userLastNameValueObject,
		Email:     userEmailValueObject,
		Password:  userPasswordValueObject,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}, nil
}
