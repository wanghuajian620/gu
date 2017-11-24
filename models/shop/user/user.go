package user

import (
	"github.com/jinzhu/gorm"
	"errors"
	"github.com/fengyfei/gu/libs/orm"
	"time"
	"github.com/fengyfei/gu/libs/security"
)

type serviceProvider struct{}

var (
	Service        *serviceProvider
	typeWechat     = "wechat"
	typePhone      = "phone"
	errLoginFailed = errors.New("invalid username or password.")
	errPassword    = errors.New("invalid password.")
)

type User struct {
	ID        int32  `gorm:"primary_key;auto_increment"`
	UserName  string `gorm:"unique;type:varchar(128)"`
	NickName  string `gorm:"type:varchar(30)"`
	Phone     string `gorm:"unique"`
	Type      string `gorm:"type:varchar(30)"`
	Pass      string `gorm:"type:varchar(128)"`
	CreatedAt *time.Time
}

func (this *serviceProvider) WechatLogin(conn orm.Connection, nickName, unionId *string) (int32, error) {

	user := &User{}
	res := &User{}
	user.UserName = *unionId
	user.NickName = *nickName
	user.Type = typeWechat

	db := conn.(*gorm.DB).Exec("USE user")

	err := db.Where("user_name = ?", *unionId).First(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// not found, create new user
			err = db.Model(&User{}).Create(&user).Error
			if err != nil {
				return 0, err
			}

			return user.ID, nil
		}
		return 0, err
	}

	return res.ID, nil
}

// register by phoneNumber
func (this *serviceProvider) PhoneRegister(conn orm.Connection, phone, password, nickName *string) error {
	salt, err := security.SaltHashGenerate(password)
	if err != nil {
		return err
	}

	now := time.Now()

	user := &User{}
	user.UserName = *phone
	user.Phone = *phone
	user.Type = typePhone
	user.NickName = *nickName
	user.Pass = string(salt)
	user.CreatedAt = &now

	db := conn.(*gorm.DB).Exec("USE shop")

	return db.Model(&User{}).Create(&user).Error
}

func (this *serviceProvider) PhoneLogin(conn orm.Connection, phone, password *string) (int32, error) {

	db := conn.(*gorm.DB).Exec("USE shop")
	user := &User{}

	err := db.Where("user_name = ?", *phone).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return 0, err
	}

	if !security.SaltHashCompare([]byte(user.Pass), password) {
		return 0, errLoginFailed
	}

	return user.ID, err
}

func (this *serviceProvider) ChangePassword(conn orm.Connection, id int32, oldPass, newPass *string) error {
	db := conn.(*gorm.DB).Exec("USE shop")
	user := &User{}

	err := db.Where("id = ?", id).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}

	if !security.SaltHashCompare([]byte(user.Pass), oldPass) {
		return errPassword
	}

	salt, err := security.SaltHashGenerate(newPass)
	if err != nil {
		return err
	}
	user.Pass = string(salt)
	return db.Save(&user).Error
}
