package model

type UserModel struct {
	BaseModel
	Email    string `gorm:"column:email"`
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}

func (UserModel) TableName() string {
	return "users"
}
