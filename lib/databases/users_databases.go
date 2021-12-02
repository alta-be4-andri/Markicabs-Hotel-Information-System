package databases

import (
	"project2/config"
	"project2/middlewares"
	"project2/models"
	"project2/plugins"
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

func GetUserByEmail(loginuser models.Users) (*models.Users, error) {
	user := models.Users{}
	tx := config.DB.Where("email=?", loginuser.Email).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	checkpass := plugins.Decrypt(loginuser.Password, user.Password)
	if !checkpass {
		return nil, nil
	}
	return &user, nil
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
func LoginUser(UserLogin models.UserLogin) (interface{}, error) {
	result := models.Get_User{}
	user := models.Users{}
	var err error
	if err = config.DB.Where("email = ?", UserLogin.Email).Find(&user).Error; err != nil {
		return nil, err
	}

	check := plugins.Decrypt(user.Password, UserLogin.Password)
	if !check {
		return 0, nil
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	config.DB.Model(user).Find(&result, user)
	return result, nil
}
