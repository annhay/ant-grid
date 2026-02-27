package utils

import (
	"ant-grid/internal/common/model"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

const (
	// PickupCodeLength 取件码长度
	PickupCodeLength = 8
	// PickupCodeChars 取件码字符集
	PickupCodeChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// GeneratePickupCode 生成取件码（全局唯一）
// 全局唯一：在同一小区内所有设备中唯一
func GeneratePickupCode(db *gorm.DB) (string, error) {
	rand.Seed(time.Now().UnixNano())
	orderModel := &model.AgOrder{}

	// 尝试生成取件码，最多尝试10次
	for i := 0; i < 10; i++ {
		code := generateRandomCode(PickupCodeLength)
		exists, err := orderModel.CheckPickupCodeExists(db, code)
		if err != nil {
			return "", err
		}
		if !exists {
			return code, nil
		}
	}

	// 如果10次都失败，使用时间戳和随机数生成一个唯一码
	timestamp := time.Now().Unix()
	random := rand.Intn(10000)
	code := formatCode(timestamp, random)
	return code, nil
}

// generateRandomCode 生成指定长度的随机码
func generateRandomCode(length int) string {
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = PickupCodeChars[rand.Intn(len(PickupCodeChars))]
	}
	return string(code)
}

// formatCode 格式化取件码
func formatCode(timestamp int64, random int) string {
	// 时间戳后4位 + 随机数4位，总共8位
	timeStr := time.Now().Format("060102150405")
	timePart := timeStr[len(timeStr)-4:]
	randomPart := timeStr[:4]
	return timePart + randomPart
}
