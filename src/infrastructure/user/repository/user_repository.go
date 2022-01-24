package repository

import (
	aggregate "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/aggregate"
	repository "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/repository"
	user_email "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/valueobject/email"
	"github.com/srcanton/go-gin-ddd-cqrs/src/infrastructure/common/persistence/model"
	"gorm.io/gorm"
)

// Repository is the repository of domain.Repository
type Repository struct {
	postgressClient *gorm.DB
}

func New(postgressClient *gorm.DB) repository.Interface {
	return &Repository{postgressClient: postgressClient}
}

func (r *Repository) Save(user *aggregate.User) error {
	userModel := model.User{
		ID:        user.ID.GetValue(),
		FirstName: user.FirstName.GetValue(),
		LastName:  user.LastName.GetValue(),
		Email:     user.Email.GetValue(),
		Password:  user.Password.GetValue(),
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return r.postgressClient.Create(&userModel).Error
}

func (r *Repository) GetUserByEmail(email user_email.Email) (*aggregate.User, error) {
	var userModel model.User
	if err := r.postgressClient.Where("email = ?", email.GetValue()).Take(&userModel).Error; err != nil {
		return nil, err
	}
	return aggregate.FromRaw(userModel.ToArray())
}
