package entity

type UserEntity struct {
	BaseEntity
	Email    string
	Name     string
	Password string
}
