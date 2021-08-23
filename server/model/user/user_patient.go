package user

type Patient struct {
	Id              int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name            string
	Birth           string
	Phone           string
	Email           string
	Address         string
	Photo           string
	AppointmentTime string
}
