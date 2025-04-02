package userhandler

type CreateUserDto struct {
	Name string `binding:"required"`
}

type UpdateUserDto struct {
	Name string `binding:"required"`
}
