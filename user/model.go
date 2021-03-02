package user

import "github.com/gofrs/uuid"

type User struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Message string `json:"message"`
}

type GetUserResponse User

type GetUserRequest struct {
	Id uuid.UUID `json:"id"`
}

type DeleteUserByIDRequest GetUserRequest
type DeleteUserByIDResponse CreateUserResponse
