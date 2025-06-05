package users

type (
	CreateUserReq struct {
		Name     string `json:"name" validate:"required,max=64"`
		Email    string `json:"email" validate:"required,email,max=64"`
		Password string `json:"password" validate:"required,min=8,max=64,regexp=^(?=.*\\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$"`
	}

	CreateUserRes struct {
		ID string `json:"id"`
	}

	LoginCredentialReq struct {
		Email    string `json:"email" validate:"required,email,max=64"`
		Password string `json:"password" validate:"required,min=8,max=64"`
	}

	LoginCredentialRes struct {
		UserId string `json:"user_id"`
		Token  string `json:"token"`
	}

	ListUserRes struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	GetUserByIdReq struct {
		ID string `json:"id" validate:"required,objectid"`
	}

	UpdateUserReq struct {
		Name  string `json:"name" validate:"max=64"`
		Email string `json:"email" validate:"email,max=64"`
	}

	// DeleteUserReq struct {
	// 	ID string `json:"id"`
	// }
)
