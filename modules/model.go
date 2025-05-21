package modules

type (
	CreateUserReq struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
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
