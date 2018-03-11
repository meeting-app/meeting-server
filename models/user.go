package models

import (
	"github.com/ezradiniz/meeting-server/database"
	"golang.org/x/crypto/bcrypt"
)

// UserModel model
type UserModel struct {
	ID           uint        `gorm:"primary_key"`
	Name         string      `gorm:"column:name"`
	Username     string      `gorm:"column:username;unique_index"`
	Email        string      `gorm:"column:email;unique_index"`
	PasswordHash string      `gorm:"column:password;not null"`
	Posts        []PostModel `gorm:"foreignkey:UserID"`
}

// UserAutoMigrate ...
func UserAutoMigrate() {
	db := database.GetDB()
	db.AutoMigrate(&UserModel{})
}

// SetPassword set user password
func (u *UserModel) SetPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *UserModel) checkPassword(password string) error {
	bytePassword1 := []byte(password)
	bytePassword2 := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(bytePassword2, bytePassword1)
}

// FindUserLogin find user by email and password
func FindUserLogin(email, password string) (UserModel, error) {
	db := database.GetDB()
	var model UserModel
	if err := db.Where("email = ?", email).First(&model).Error; err != nil {
		return model, err
	}
	return model, model.checkPassword(password)
}

// AddUser insert user into database
func AddUser(user interface{}) error {
	db := database.GetDB()
	return db.Save(user).Error
}

// FindUserByEmail find user by email
func FindUserByEmail(email string) (UserModel, error) {
	db := database.GetDB()
	var model UserModel
	return model, db.Where("email = ?", email).First(&model).Error
}
