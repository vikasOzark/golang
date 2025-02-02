package models

import (
	"database/sql"
	"time"

	valid "github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uint         `gorm:"primaryKey" json:"id" valid:"optional"`
	Name        string       `gorm:"not null" json:"name" valid:"required~Name is required field."`
	Email       *string      `gorm:"unique;not null" json:"email" valid:"email,required~Email is required field."`
	ActivatedAt sql.NullTime `gorm:"type:timestamp"`
	CompanyId   string       `gorm:"allow null" json:"company_id"`
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime"`
	Password    string       `gorm:"not null" json:"-" valid:"required~Password can not be empty."`
}

func (user *User) Validate() error {
	_, err := valid.ValidateStruct(user)
	return err
}

func (User) TableName() string {
	return "users"
}

func (user *User) EncyPassword() {

	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
}

func (user *User) Compare(password string) bool {
	hashedPassword := []byte(user.Password)
	passwordBytes := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashedPassword, passwordBytes)
	return err == nil
}
