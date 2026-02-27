package model

import (
	"time"

	"gorm.io/gorm"
)

type AgOrder struct {
	Id          uint64    `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	DeviceId    int64     `gorm:"column:device_id;type:bigint;comment:设备ID;not null;" json:"device_id"`                    // 设备 ID
	OrderCode   string    `gorm:"column:order_code;type:varchar(30);comment:快递单号;not null;" json:"order_code"`             // 快递单号
	Phone       string    `gorm:"column:phone;type:varchar(50);comment:用户手机号;not null;" json:"phone"`                      // 用户手机号
	PickupCode  string    `gorm:"column:pickup_code;type:char(8);comment:取件码（局部唯一）;not null;" json:"pickup_code"`          // 取件码（局部唯一）
	Status      string    `gorm:"column:status;type:varchar(10);comment:状态：0待取 1已取 2超时;not null;default:0;" json:"status"` // 状态：0待取 1已取 2超时
	ExpireTime  time.Time `gorm:"column:expire_time;type:datetime(3);comment:超时时间;not null;" json:"expire_time"`           // 超时时间
	OvertimeFee float64   `gorm:"column:overtime_fee;type:decimal(10, 2);comment:超时计费;default:0.00;" json:"overtime_fee"`  // 超时计费
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
}

func (ao *AgOrder) TableName() string {
	return "ag_order"
}
func (ao *AgOrder) GetOrderByDeviceAndCode(db *gorm.DB, deviceId int64, code string) error {
	return db.Where("device_id = ? and pickup_code = ?", deviceId, code).First(&ao).Error
}
func (ao *AgOrder) Editor(db *gorm.DB) error {
	return db.Updates(&ao).Error
}

// CheckPickupCodeExists 检查取件码是否已存在（在所有设备中）
func (ao *AgOrder) CheckPickupCodeExists(db *gorm.DB, pickupCode string) (bool, error) {
	var count int64
	err := db.Model(&AgOrder{}).Where("pickup_code = ?", pickupCode).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CreateOrder 创建订单
func (ao *AgOrder) CreateOrder(db *gorm.DB) error {
	return db.Create(ao).Error
}
