package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string // 昵称
	LoginPwd  string // 密码
	Email     string // 邮箱
	Phone     string // 电话
	Flag      string // 启用标志(1-启用 0-停用)
	AvatarUrl string // 头像
	Token     string
	TokenDate time.Time // Token时间
}

// UserRefreshManagerToken
//
//	@Description:	修改指定用户的token数据
//	@param			token	数据格式	<token_value:timestamp>
//	@return			err
func UserRefreshManagerToken(db *gorm.DB, userId int64, token string) (err error) {
	res := db.Model(&User{}).Where("id = ?", userId).Update("manager_token", token)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

// UserSelectIdByToken token查询用户数据 token = "HASH"
func UserSelectIdByToken(db *gorm.DB, token string) (user User, err error) {
	err = db.Table("users").
		Where("token LIKE ? AND flag = ?", token+":%", "1").Take(&user).Error
	return
}

// UserRefreshToken
//
//	@Description:	修改指定用户的token数据
//	@param			token	数据格式	<token_value:timestamp>
//	@return			err
func UserRefreshToken(db *gorm.DB, userId int64, token string) (err error) {
	res := db.Model(&User{}).Where("id = ?", userId).Update("token", token)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected == 0")
	}
	return nil
}

func UserSelectIdByManagerToken(db *gorm.DB, token string) (user User, err error) {
	err = db.Table("users").
		Where("manager_token LIKE ? AND flag = ?", token+":%", "1").Take(&user).Error
	return
}

// SelectUserByPhone 根据手机号查询用户
func SelectUserByPhone(db *gorm.DB, phone string) (*User, error) {
	user := &User{}
	err := db.Model(&User{}).Where("phone = ?", phone).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// SelectUserByEmail 根据 邮箱查询用户
func SelectUserByEmail(db *gorm.DB, email string) (*User, error) {
	user := &User{}
	err := db.Model(&User{}).Where("email = ?", email).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

func (f *User) Update(db *gorm.DB) (err error) {
	res := db.Model(&f).Select("*").Updates(f)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("user Update res.RowsAffected == 0")
	}
	return nil
}

// SelectUserByUserId 根据用户Id查询用户
func SelectUserByUserId(db *gorm.DB, userId int64) (user User, err error) {
	if err := db.Model(&User{}).Where("id = ?", userId).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
