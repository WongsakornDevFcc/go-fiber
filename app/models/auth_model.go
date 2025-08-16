package models

// SignUp struct to describe register a new user.
type SignUp struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
	UserRole string `json:"user_role" validate:"required,lte=25"`
}

type Signin1 struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// SignIn struct to describe login user.
type SignIn struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}
