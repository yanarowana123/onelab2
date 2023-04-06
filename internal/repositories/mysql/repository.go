package repository

import (
	"database/sql"
	"github.com/yanarowana123/onelab2/internal/models"
)

type MysqlUserRepository struct {
	db *sql.DB
}

// Если используешь ДБ то нужно сделать в отдельном файле функцию подключения
func NewMysqlUserRepository(db *sql.DB) *MysqlUserRepository {
	return &MysqlUserRepository{
		db,
	}
}

func (r *MysqlUserRepository) CreateUser(user models.CreateUserReq) (*models.UserResponse, error) {
	_, err := r.db.Exec("insert into users (name, login, password) values (?,?,?)",
		user.Name, user.Login, user.Password)

	if err != nil {
		return nil, err
	}

	return models.NewUserResponse(user.Name, user.Login), nil
}

func (r *MysqlUserRepository) GetUser(login string) (*models.UserResponse, error) {
	var userResponse models.UserResponse
	err := r.db.QueryRow("SELECT login, name FROM users WHERE login=?", login).Scan(&userResponse.Login, &userResponse.Name)

	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}
