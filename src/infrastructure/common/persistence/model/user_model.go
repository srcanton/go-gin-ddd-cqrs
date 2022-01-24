package model

// User is the model of users
type User struct {
	ID        string  `gorm:"primary_key;column:id;type:uuid" json:"id"`
	FirstName string  `gorm:"size:100;not null;" json:"first_name"`
	LastName  string  `gorm:"size:100;not null;" json:"last_name"`
	Email     string  `gorm:"size:100;not null;unique" json:"email"`
	Password  string  `gorm:"size:100;not null;" json:"password"`
	CreatedAt string  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt string  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

func (u *User) ToArray() map[string]interface{} {
	user := map[string]interface{}{
		"ID":        u.ID,
		"FirstName": u.FirstName,
		"LastName":  u.LastName,
		"Email":     u.Email,
		"Password":  u.Password,
		"CreatedAt": u.CreatedAt,
		"UpdatedAt": u.UpdatedAt,
	}
	return user
}
