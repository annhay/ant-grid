package server

import (
	"ant-grid/internal/common/global"
	"ant-grid/internal/common/model"
	"ant-grid/internal/common/utils"
	"ant-grid/internal/proto"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	proto.UnimplementedOrderServer
}

// CreateOrder 创建订单（快递员存件）
func (s *Server) CreateOrder(_ context.Context, in *proto.CreateOrderReq) (*proto.CreateOrderResp, error) {
	// 生成全局唯一的取件码
	pickupCode, err := utils.GeneratePickupCode(global.DB)
	if err != nil {
		return nil, err
	}

	// 创建订单
	order := model.AgOrder{
		DeviceId:   in.DeviceId,
		OrderCode:  in.OrderCode,
		Phone:      in.Phone,
		PickupCode: pickupCode,
		Status:     "0",                            // 0待取
		ExpireTime: time.Now().Add(48 * time.Hour), // 48小时后超时
	}

	if err := order.CreateOrder(global.DB); err != nil {
		return nil, err
	}

	resp := &proto.CreateOrderResp{
		PickupCode: pickupCode,
		OrderCode:  in.OrderCode,
	}
	return resp, nil
}

// PickUp 取件
func (s *Server) PickUp(_ context.Context, in *proto.PickUpReq) (*proto.PickUpResp, error) {
	var order model.AgOrder
	if err := order.GetOrderByDeviceAndCode(global.DB, in.DeviceId, in.PickupCode); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("取件码错误")
		}
	}

	// 检查是否超时，如果超时计算最终超时费用
	if order.Status == "0" && time.Now().After(order.ExpireTime) {
		// 计算超时费用（从过期时间开始计算，每超时1天收费1元）
		overtimeDays := time.Since(order.ExpireTime).Hours() / 24
		if overtimeDays > 0 {
			order.OvertimeFee = float64(int(overtimeDays)) * 1.0
		}
	}

	order.Status = "1" //订单状态改为1已取件
	if err := order.Editor(global.DB); err != nil {
		return nil, err
	}

	resp := &proto.PickUpResp{
		OvertimeFee: order.OvertimeFee,
	}
	return resp, nil
}
