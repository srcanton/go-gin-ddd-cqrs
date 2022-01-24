package get_user_by_email

import aggregate "github.com/srcanton/go-gin-ddd-cqrs/src/domain/user/aggregate"

type GetUserByIdQueryResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func NewGetUserByIdQueryResponse(user aggregate.User) *GetUserByIdQueryResponse {
	return &GetUserByIdQueryResponse{
		ID:    user.ID.GetValue(),
		Email: user.Email.GetValue(),
	}
}
