package model

import "time"

type AgDevice struct { //设备信息表
	Id            uint64    `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	DeviceName    string    `gorm:"column:device_name;type:varchar(100);comment:设备名称;not null;" json:"device_name"`       // 设备名称
	DeviceAddress string    `gorm:"column:device_address;type:varchar(255);comment:设备地址;not null;" json:"device_address"` // 设备地址
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
}

func (ad *AgDevice) TableName() string {
	return "ag_device"
}
