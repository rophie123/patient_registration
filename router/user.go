package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rophie123/patient_registration/model"
	"github.com/rophie123/patient_registration/utils"
	"path/filepath"
)

type LoginForm struct {
	UserName string
	Password string
}

type RegForm struct {
	Name            string
	Birth           int
	Phone           string
	Email           string
	Address         string
	Photo           string
	AppointmentTime int
}

func Login(c *gin.Context) {
	var form LoginForm
	if err := c.Bind(&form); err != nil {
		ReturnError(c, "parameter error")
		return
	}
	count := 0
	model.DB.Table("admins").Where("user_name=? and password=?", form.UserName, form.Password).Count(&count)
	if count == 0 {
		ReturnError(c, "login error")
		return
	}
	token := utils.NewToken(form.UserName)
	ReturnData(c, token)
	return
}

func Info(c *gin.Context) {
	userName := utils.ParseToken(c.GetHeader("X-Token"))
	ReturnData(c, gin.H{
		"roles": []string{"admin"},
		"name":  userName,
	})
}

func Reg(c *gin.Context) {
	var form model.Patient
	if err := c.Bind(&form); err != nil {
		ReturnError(c, err.Error())
		return
	}
	if err := model.DB.FirstOrCreate(&form, model.Patient{Phone: form.Phone}).Error; err != nil {
		ReturnError(c, err.Error())
		return
	}
	ReturnOK(c, "save success")
}

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		ReturnError(c, "upload error")
		return
	}
	var dest = filepath.Join("upload", file.Filename)
	err = c.SaveUploadedFile(file, dest)
	if err != nil {
		ReturnError(c, "upload error")
		return
	}
	ReturnData(c, file.Filename)
	return
}

func List(c *gin.Context) {
	var patients []model.Patient
	model.DB.Find(&patients)
	ReturnData(c, patients)
}
