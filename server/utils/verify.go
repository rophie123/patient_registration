package utils

var (
	LoginVerify = Rules{
		"UserName": {NotEmpty()},
		"Password": {NotEmpty()},
	}
	RegisterVerify = Rules{
		"Name":            {NotEmpty()},
		"Birth":           {NotEmpty()},
		"Phone":           {NotEmpty(), RegexpMatch("^(13[0-9]|14[5|7]|15[0|1|2|3|4|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\\d{8}$")},
		"Email":           {NotEmpty(), RegexpMatch("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$")},
		"Address":         {NotEmpty()},
		"Photo":           {NotEmpty()},
		"AppointmentTime": {NotEmpty()},
	}
)
