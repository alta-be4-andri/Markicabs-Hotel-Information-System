package databases

import (
	"project2/config"
	"project2/middlewares"
	"project2/models"

	"golang.org/x/crypto/bcrypt"
)

var user models.Users

// function database untuk menambahkan user baru (registrasi)
func CreateUser(user *models.Users) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// function database untuk menampilkan user by id
func GetUser(id int) (interface{}, error) {
	result := models.Get_User{}
	err := config.DB.Model(user).Find(&result, id)
	rows_affected := err.RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	return result, nil
}

// function database untuk memperbarui data user by id
func UpdateUser(id int, user *models.Users) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return nil, err
	}
	config.DB.First(&user, id)
	return user, nil
}

// function database untuk menghapus data user by id
func DeleteUser(id int) (interface{}, error) {
	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// function database untuk melakukan login
func LoginUsers(user models.UserLogin) (string, error) {
	var err error
	userpassword := models.Users{}
	if err = config.DB.Where("email = ?", user.Email).First(&userpassword).Error; err != nil {
		return "", err
	}
	hashpassword, _ := GeneratehashPassword(userpassword.Password)
	check := CheckPasswordHash(user.Password, hashpassword)
	if !check {
		return "", nil
	}
	userpassword.Password = hashpassword

	userpassword.Token, err = middlewares.CreateToken(int(userpassword.ID))
	if err != nil {
		return "", err
	}

	if err := config.DB.Save(&userpassword).Error; err != nil {
		return "", err
	}
	return userpassword.Token, nil
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
