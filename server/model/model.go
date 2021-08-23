package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rophie123/patient_registration/server/model/user"
	"github.com/rophie123/patient_registration/server/utils"
)

var DB *gorm.DB

func Init() (err error) {
	DB, err = gorm.Open("sqlite3", fmt.Sprintf("%s?loc=Asia/Shanghai", "database.db"))
	if err != nil {
		return
	}
	DB.LogMode(false)
	DB.AutoMigrate(user.Admin{}, user.Patient{})
	count := 0
	defUser := "admin"
	defPwd := "123456"
	DB.Table("admin").Where("user_name = ?", defUser).Count(&count)
	if count == 0 {
		DB.Table("admins").Create(&user.Admin{
			UserName: defUser,
			Password: utils.MD5(defPwd),
		})
	}
	return
}
