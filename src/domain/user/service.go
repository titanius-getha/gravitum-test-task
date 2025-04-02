package user

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) (*UserService, error) {
	return &UserService{repo}, repo.Migrate()
}

func (u *UserService) GetByID(ID uint) (User, error) {
	user, err := u.repo.GetByID(ID)
	return user, err
}

func (u *UserService) Create(name string) (User, error) {
	newUser := User{Name: name}
	return u.repo.Save(newUser)
}

func (u *UserService) Update(ID uint, name string) (User, error) {
	user, err := u.GetByID(ID)
	if err != nil {
		return user, err
	}

	user.Name = name

	return u.repo.Save(user)
}
