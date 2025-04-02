package user

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("пользователь не найден")
)

type UserRepository interface {
	// Автоматическая миграциия данных (для простоты)
	Migrate() error
	// Получить пользователя по ID.
	// Вернет ошибку ErrUserNotFound, если пользователя не существует.
	GetByID(ID uint) (User, error)
	// Сохранить пользователя. Если у пользователя нет поля ID, создает нового пользователя.
	// Возвращает созданного или обновленного пользователя.
	Save(user User) (User, error)
}

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) UserRepository {
	return &PostgresRepository{db}
}

func (u *PostgresRepository) Migrate() error {
	return u.db.AutoMigrate(&User{})
}

func (u *PostgresRepository) GetByID(ID uint) (User, error) {
	var user User
	err := u.db.Session(&gorm.Session{}).Where("id=?", ID).Take(&user).Error

	if errors.Is(gorm.ErrRecordNotFound, err) {
		return user, ErrUserNotFound
	}

	return user, err
}

func (u *PostgresRepository) Save(user User) (User, error) {
	err := u.db.Session(&gorm.Session{}).Save(&user).Error
	return user, err
}
