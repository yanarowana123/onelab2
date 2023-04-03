package services

type Manager struct {
	User IUserService
}

func NewManager(userService IUserService) *Manager {
	return &Manager{User: userService}
}
