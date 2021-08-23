package service

import (
	"errors"
	"github.com/rophie123/patient_registration/server/model"
	"github.com/rophie123/patient_registration/server/model/user"
)

type UserService struct {
}

func (userService *UserService) Login(u *user.Admin) error {
	count := 0
	err := model.DB.Table("admins").Where("user_name=? and password=?", u.UserName, u.Password).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("login error")
	}
	return nil
}

func (userService *UserService) Patients() (err error, patients []user.Patient) {
	err = model.DB.Find(&patients).Error
	return
}

func (userService *UserService) NewPatient(p *user.Patient) error {
	count := 0
	model.DB.Table("patients").Where("phone=?", p.Phone).Count(&count)
	if count > 0 {
		return errors.New("patient already exists")
	}
	return model.DB.FirstOrCreate(p, user.Patient{Phone: p.Phone}).Error
}
