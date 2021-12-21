package main

import (
	"context"
	"gotest/kitex_gen/api"
	"time"
)

type Order struct {
	ID        int64
	RoomID    int64
	UserID    int64
	Time      time.Time
	ProductID int64
	Price     int64
}

// GotestImpl implements the last service interface defined in the IDL.
type GotestImpl struct{}

// Add implements the GotestImpl interface.
func (s *GotestImpl) Add(ctx context.Context, req *api.AddRequest) (resp *api.AddResponse, err error) {
	// TODO: Your code here...
	order := Order{ID: req.ID, RoomID: req.RoomID, UserID: req.UserID, Price: req.Price, ProductID: req.ProductID}
	order.Time, err = time.Parse(time.ANSIC, req.Time)
	if err != nil {
		panic(err)
	}
	if err := DbClient.DB.Create(order).Error; err != nil {
		panic(err)
	}
	resp = &api.AddResponse{Message: "插入成功"}
	return
}

// Delete implements the GotestImpl interface.
func (s *GotestImpl) Delete(ctx context.Context, req *api.DeleteRequest) (resp *api.DeleteResponse, err error) {
	// TODO: Your code here...

	if err := DbClient.DB.Delete(&Order{}, req.ID).Error; err != nil {
		panic(err)
	}
	resp = &api.DeleteResponse{Message: "删除成功"}
	return
}

// Select implements the GotestImpl interface.
func (s *GotestImpl) Select(ctx context.Context, req *api.SelectRequest) (resp *api.SelectResponse, err error) {
	// TODO: Your code here...
	var sum int64
	if err := DbClient.DB.Table("orders").Select("ifnull(sum(Price),0)").Where("room_id = ?", req.RoomID).Find(&sum).Error; err != nil {
		panic(err)
	}
	resp = &api.SelectResponse{Price: sum, Message: ""}
	return
}
