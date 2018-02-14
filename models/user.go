package users

import "golang.org/x/crypto/bcrypt"

type UserModel struct {
	ID           uint   `gorm:"primary_key"`
	Username     string `gorm:"column:username"`
	Email        string `gorm:"column:email;unique_index"`
	PasswordHash string `gorm:"column:password;not null"`
}

func (u *UserModel) setPassword(password string) error {
	if len(password) < 6 {
		return error.New("invalid password")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *UserModel) checkPassword(password string) error {
	bytePassword1 := []byte(password)
	bytePassword2 := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(bytePassword1, bytePassword2)
}
