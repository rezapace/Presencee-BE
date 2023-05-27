package payload

type (
	CreateUserRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=5"`
		Role     string `json:"role" form:"role" validate:"required"`
	}

	CreateUserResponse struct {
		UserID uint   `json:"user_id"`
		Role   string `json:"role"`
		Token  string `json:"token"`
	}

	LoginUserRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=5"`
	}

	LoginUserResponse struct {
		UserID uint   `json:"user_id"`
		Token  string `json:"token"`
	}
)
