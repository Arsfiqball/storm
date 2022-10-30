package user

type UserEntity struct {
	ID       ID       `gorm:"primaryKey;column:id"`
	Email    Email    `gorm:"column:email"`
	Password Password `gorm:"column:password"`
}

func (UserEntity) TableName() string {
	return "user"
}
