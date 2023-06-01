package entity

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	*gorm.Model

	ID           uint      `db:"id"`
	BuyerEmail   string    `db:"buyer_email"`
	BuyerAddress string    `db:"buyer_address"`
	ProductID    uint      `db:"product_id"`
	OrderTime    time.Time `db:"order_date"`
	Product      Product   //diambil dari struct Product
}
