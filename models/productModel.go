package models

import "time"

type Product struct {
	Id               int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT; primaryKey: true"`
	Sku              string    `json:"sku"`
	Name             string    `json:"name"`
	Stock            int       `json:"stock"`
	Price            int       `json:"price"`
	Image            string    `json:"image"`
	TotalFinalPrice  int       `"total_final_price"`
	TotalNormalPrice int       `"total_normal_price"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Category         int       `json:"category"`
	DiscountId       int       `json:"discount_id"`
}
