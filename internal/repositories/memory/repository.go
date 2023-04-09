package memory

//
//type MemoryUserRepository struct {
//	users map[string]models.CreateUserReq
//	lock  *sync.RWMutex
//}
//
//func NewMemoryUserRepository() *MemoryUserRepository {
//	usersMap := make(map[string]models.CreateUserReq)
//	return &MemoryUserRepository{
//		users: usersMap,
//		lock:  &sync.RWMutex{},
//	}
//}
//
//func (r *MemoryUserRepository) CreateUser(user models.CreateUserReq) (*models.UserResponse, error) {
//	userResponse := r.getUser(user.Login)
//
//	if userResponse != nil {
//		return userResponse, nil
//	}
//
//	r.createUser(user)
//
//	return models.NewUserResponse(user.Name, user.Login), nil
//}
//
//func (r *MemoryUserRepository) GetUser(login string) (*models.UserResponse, error) {
//	userResponse := r.getUser(login)
//
//	if userResponse != nil {
//		return userResponse, nil
//	}
//
//	return nil, nil
//}
//
//func (r *MemoryUserRepository) getUser(login string) *models.UserResponse {
//	r.lock.RLock()
//	defer r.lock.RUnlock()
//
//	if user, ok := r.users[login]; ok {
//		return models.NewUserResponse(user.Name, user.Login)
//	}
//
//	return nil
//}
//
//func (r *MemoryUserRepository) createUser(user models.CreateUserReq) {
//	r.lock.Lock()
//	defer r.lock.Unlock()
//
//	r.users[user.Login] = user
//}
