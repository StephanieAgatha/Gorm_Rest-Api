package entity

import "gorm.io/gorm"

type Product struct {

	//generate database using gorm model
	*gorm.Model

	ID          uint   `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       int64  `db:"price"`
}
