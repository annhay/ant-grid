package scheduler

import (
	"ant-grid/internal/common/global"
	"ant-grid/internal/common/model"
	"log"
	"time"
)

// StartOrderScheduler 启动订单定时任务
func StartOrderScheduler() {
	go func() {
		ticker := time.NewTicker(5 * time.Minute) // 每5分钟检查一次
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				checkExpiredOrders()
			}
		}
	}()
	log.Println("订单定时任务已启动")
}

// checkExpiredOrders 检查过期订单
func checkExpiredOrders() {
	var orders []model.AgOrder
	// 查找所有状态为待取且已超时的订单
	result := global.DB.Where("status = ? AND expire_time < ?", "0", time.Now()).Find(&orders)
	if result.Error != nil {
		log.Printf("查询过期订单失败: %v\n", result.Error)
		return
	}

	for _, order := range orders {
		// 更新订单状态为超时
		order.Status = "2"
		// 计算超时费用（从过期时间开始计算，每超时1天收费1元）
		overtimeDays := time.Since(order.ExpireTime).Hours() / 24
		if overtimeDays > 0 {
			order.OvertimeFee = float64(int(overtimeDays)) * 1.0
		}
		
		if err := order.Editor(global.DB); err != nil {
			log.Printf("更新订单状态失败 [订单ID: %d]: %v\n", order.Id, err)
		} else {
			log.Printf("订单已标记为超时 [订单ID: %d, 超时费用: %.2f]\n", order.Id, order.OvertimeFee)
		}
	}
}
