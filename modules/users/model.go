package users

type (
	CreateUserReq struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginCredentialReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
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
		ID string `json:"id"`
	}

	UpdateUserReq struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	DeleteUserReq struct {
		ID string `json:"id"`
	}
)
