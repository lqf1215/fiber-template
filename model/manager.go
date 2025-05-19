package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

// Manager struct
type Manager struct {
	gorm.Model
	Username  string
	Token     string
	TokenDate time.Time // Token时间
	Password  string
}

func GetManagerByUsername(db *gorm.DB, username string) (*Manager, error) {
	m := Manager{}
	if err := db.Model(&m).Where("username = ?", username).Take(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *Manager) UpdateManager(db *gorm.DB) error {
	res := db.Model(&m).Updates(m)
	if res.RowsAffected > 0 {
		return errors.New("res.RowsAffected > 0")
	}
	return nil
}

// 新增用户
func (m *Manager) InsertManager(db *gorm.DB) error {
	return db.Create(m).Error
}

// ManagerSelectIdByToken token查询用户数据 token = "HASH"
func ManagerSelectIdByToken(db *gorm.DB, token string) (manager Manager, err error) {
	err = db.Table("managers").
		Where("token LIKE ?", token+":%").Take(&manager).Error
	return
}
func ManagerRefreshToken(db *gorm.DB, userId int64, token string) (err error) {
	res := db.Model(&Manager{}).Where("id = ?", userId).Updates(map[string]interface{}{
		"token":      token,
		"token_date": time.Now(),
	})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("res.RowsAffected ==0")
	}
	return nil
}
