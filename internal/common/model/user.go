package model

import (
	"time"

	"gorm.io/gorm"
)

type AgUser struct { //用户信息表
	Id            uint64    `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	Phone         string    `gorm:"column:phone;type:varchar(50);comment:手机号;not null;" json:"phone"`                            // 手机号
	Password      string    `gorm:"column:password;type:varchar(64);comment:密码;not null;" json:"password"`                       // 密码
	Cover         string    `gorm:"column:cover;type:varchar(255);comment:头像;default:NULL;" json:"cover"`                        // 头像
	Nickname      string    `gorm:"column:nickname;type:varchar(50);comment:昵称;default:NULL;" json:"nickname"`                   // 昵称
	RealName      string    `gorm:"column:real_name;type:varchar(100);comment:真实姓名;default:NULL;" json:"real_name"`              // 真实姓名
	IdCard        string    `gorm:"column:id_card;type:varchar(50);comment:身份证号;default:NULL;" json:"id_card"`                   // 身份证号
	LastLoginTime time.Time `gorm:"column:last_login_time;type:datetime(3);comment:最后登录时间;default:NULL;" json:"last_login_time"` // 最后登录时间
	Status        string    `gorm:"column:status;type:varchar(10);comment:状态：1正常 2注销;not null;default:1;" json:"status"`         // 状态：1正常 2注销
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
}

func (au *AgUser) TableName() string {
	return "ag_user"
}
func (au *AgUser) GetUserByPhone(db *gorm.DB, phone string) error {
	return db.Where("phone = ?", phone).First(&au).Error
}
func (au *AgUser) CreateUser(db *gorm.DB) error {
	return db.Create(&au).Error
}
func (au *AgUser) Editor(db *gorm.DB) error {
	return db.Updates(&au).Error
}
func (au *AgUser) GetUserById(db *gorm.DB, id int) error {
	return db.Where("id = ?", id).First(&au).Error
}
