package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductId uint32 `gorm:"type:int(11);not null"`
	Qty       uint32 `gorm:"type:int(11);not null"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, c *Cart) error {
	var rows Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).
		First(&rows).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if rows.ID > 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).
			UpdateColumn("qty", gorm.Expr("qty+?", c.Qty)).Error
	}

	return db.WithContext(ctx).Create(c).Error
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userId).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint32) ([]*Cart, error) {
	var rows []*Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: userId}).Find(&rows).Error
	return rows, err
}
