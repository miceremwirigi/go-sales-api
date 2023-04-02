package models

import "time"

type Discount struct {
	Id              int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT; primaryKey: true`
	Qty             int       `json:"qty"`
	Type            string    `json:"type"`
	Result          string    `json:"result"`
	ExpiredAt       int       `json:"expired"`
	AxpiredAtFormat string    `json:"axpiredAtFormat"`
	StringFormat    string    `json:"stringFormat"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
