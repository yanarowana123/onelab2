package models

type CreateUserReq struct {
	Name     string
	Login    string
	Password string
}

type UserResponse struct {
	Name  string
	Login string
}

// мы не используем конструкторы для простых моделей
func NewCreateUserReq(name, login, password string) *CreateUserReq {
	//TODO ADD VALIDATION
	return &CreateUserReq{
		Name:     name,
		Login:    login,
		Password: password,
	}
}

func NewUserResponse(name, login string) *UserResponse {
	return &UserResponse{name, login}
}
