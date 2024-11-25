package model

import (
	"context"
	"gorm.io/gorm"
)

type Consignee struct {
	Email    string
	Street   string
	City     string
	Province string
	Country  string
	ZipCode  int32
}

type Order struct {
	gorm.Model
	OrderId      string      `gorm:"type:varchar(100);uniqueIndex"`
	UserId       uint32      `gorm:"type:int(11)"`
	UserCurrency string      `gorm:"type:varchar(10)"`
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer:references:OrderId"`
}

func (Order) TableName() string {
	return "order"
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	err := db.WithContext(ctx).Model(&Order{}).Where("user_id = ?", userId).Preload("OrderItem").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
