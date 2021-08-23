package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rophie123/patient_registration/server/model/common/response"
	"github.com/rophie123/patient_registration/server/model/user"
	"github.com/rophie123/patient_registration/server/model/user/request"
	userResponse "github.com/rophie123/patient_registration/server/model/user/response"
	"github.com/rophie123/patient_registration/server/service"
	"github.com/rophie123/patient_registration/server/utils"
	"path/filepath"
)

type UserApiHandler struct {
}

func (u *UserApiHandler) Login(c *gin.Context) {
	var form request.LoginForm
	if err := c.Bind(&form); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	if err := utils.Verify(form, utils.LoginVerify); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	s := service.UserService{}
	if err := s.Login(&user.Admin{UserName: form.UserName, Password: form.Password}); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	token := utils.NewToken(form.UserName)
	response.OkWithData(c, token)
	return
}

func (u *UserApiHandler) Info(c *gin.Context) {
	userName := utils.ParseToken(c.GetHeader("X-Token"))
	response.OkWithData(c, userResponse.UserInfoResponse{Name: userName, Roles: []string{"admin"}})
}

func (u *UserApiHandler) Reg(c *gin.Context) {
	var form request.RegForm
	if err := c.Bind(&form); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	if err := utils.Verify(form, utils.RegisterVerify); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	patient := &user.Patient{
		Name:            form.Name,
		Birth:           form.Birth,
		Phone:           form.Phone,
		Email:           form.Email,
		Address:         form.Address,
		Photo:           form.Photo,
		AppointmentTime: form.AppointmentTime,
	}
	s := service.UserService{}
	if err := s.NewPatient(patient); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithMessage(c, "save success")
}

func (u *UserApiHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage(c, "upload error")
		return
	}
	var dest = filepath.Join("upload", file.Filename)
	err = c.SaveUploadedFile(file, dest)
	if err != nil {
		response.FailWithMessage(c, "upload error")
		return
	}
	response.OkWithData(c, file.Filename)
	return
}

func (u *UserApiHandler) List(c *gin.Context) {
	s := service.UserService{}
	err, patients := s.Patients()
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, patients)
}
