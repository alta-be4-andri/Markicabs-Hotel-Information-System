package databases

import (
	"project2/config"
	"project2/models"
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

/*
// function login database untuk mendapatkan token
func LoginUser(plan_pass string, user *models.Users) (interface{}, error) {
	err := config.DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return nil, err
	}

	// cek plan password dengan hash password
	match := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plan_pass))
	if match != nil {
		return nil, match
	}
	user.Token, err = middlewares.CreateToken(int(user.ID)) // generate token
	if err != nil {
		return nil, err
	}
	if err = config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user.Token, nil
}
*/
