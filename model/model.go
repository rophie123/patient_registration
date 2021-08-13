package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rophie123/patient_registration/utils"
)

type Admin struct {
	Id       int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	UserName string
	Password string
}

type Patient struct {
	Id              int `gorm:"PRIMARY_KEY;AUTO_INCREMENT" form:"-"`
	Name            string
	Birth           string
	Phone           string
	Email           string
	Address         string
	Photo           string
	AppointmentTime string `json:"appointment_time"`
}

var DB *gorm.DB

func Init() (err error) {
	DB, err = gorm.Open("sqlite3", fmt.Sprintf("%s?loc=Asia/Shanghai", "database.db"))
	if err != nil {
		return
	}
	DB.LogMode(false)
	DB.AutoMigrate(Admin{}, Patient{})
	count := 0
	defUser := "admin"
	defPwd := "123456"
	DB.Table("admin").Where("user_name = ?", defUser).Count(&count)
	if count == 0 {
		DB.Table("admins").Create(&Admin{
			UserName: defUser,
			Password: utils.MD5(defPwd),
		})
	}
	return
}
