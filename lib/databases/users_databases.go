package databases

import (
	"project2/config"
	"project2/crypt"
	"project2/middlewares"
	"project2/models"
)

var user models.Users

// function database untuk menambahkan user baru (registrasi)
func CreateUser(user *models.Users) (*models.Users, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// function database untuk menampilkan user by id
func GetUser(id int) (interface{}, error) {
	users := models.Users{}
	type get_user struct {
		Nama  string
		Email string
	}
	err := config.DB.Find(&users, id)
	rows_affected := config.DB.Find(&users, id).RowsAffected
	if err.Error != nil || rows_affected < 1 {
		return nil, err.Error
	}
	return get_user{users.Nama, users.Email}, nil
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
	config.DB.First(&user, id)
	return user, nil
}

// function database untuk melakukan login
func LoginUser(UserLogin models.UserLogin) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ?", UserLogin.Email).First(&user).Error; err != nil {
		return nil, err
	}

	check := crypt.Decrypt(user.Password, UserLogin.Password)
	if !check {
		return nil, nil
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user.Token, nil
}
