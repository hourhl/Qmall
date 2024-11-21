package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type PaymentLog struct {
	gorm.Model
	UserId        uint32    `json:"userId"`
	OrderId       string    `json:"OrderId"`
	TransactionId string    `json:"TransactionId"`
	Amout         float32   `json:"amout"`
	PayAt         time.Time `json:"payAt"`
}

func (PaymentLog) TableName() string {
	return "paymentLog"
}

func CreatePaymentLog(ctx context.Context, db *gorm.DB, paymentLog *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(paymentLog).Error
}
