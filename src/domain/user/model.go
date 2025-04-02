package user

type User struct {
	ID   uint   `gorm:"primary;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
}

func (User) TableName() string {
	return "users"
}
