package models

import "time"

type Order struct {
	Id            int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT; primaryKey: true"`
	CashierID     int       `json:"cashier_id"`
	PaymentTypeId int       `json:"paymentsTypeId"`
	TotalPrice    int       `json:"totalPrice"`
	TotalPaid     int       `json:"totalPaid"`
	TotalReturn   int       `json:"totalReturn"`
	ReceiptId     string    `json:"receiptId"`
	IsDownload    int       `json:"is_download"`
	ProductId     string    `json:"productId"`
	Quantities    string    `json:"quantities"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
