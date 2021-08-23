package user

type Admin struct {
	Id       int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	UserName string
	Password string
}
