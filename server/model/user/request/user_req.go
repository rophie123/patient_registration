package request

type LoginForm struct {
	UserName string
	Password string
}

type RegForm struct {
	Name            string
	Birth           string
	Phone           string
	Email           string
	Address         string
	Photo           string
	AppointmentTime string `json:"appointment_time"`
}
