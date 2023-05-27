package database

import (
	"presensee_project/config"
	"presensee_project/model"
)

func CreateUser(user *model.User) error {
	err := config.DB.Create(&user).Error

	if err != nil {
		return err
	}
	return nil
}

func GetUsers() (users []model.User, err error) {
	if err = config.DB.Model(&model.User{}).Preload("Blogs").Find(&users).Error; err != nil {
		return
	}
	return
}

func GetUser(user *model.User) (err error) {
	if err = config.DB.First(&user).Error; err != nil {
		return
	}
	return
}

func GetUserWithBlog(id uint) (user model.User, err error) {
	user.ID = id
	if err = config.DB.Model(&model.User{}).Preload("Blogs").First(&user).Error; err != nil {
		return
	}
	return
}

func UpdateUser(user *model.User) error {
	err := config.DB.Updates(&user).Error

	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *model.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(requestUser *model.User) (model.User, error) {
	var user model.User

	err := config.DB.Where(&model.User{Email: requestUser.Email}).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
