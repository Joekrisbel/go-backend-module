package models

import "time"

type Participant struct {
	Id                       uint      `json:"id;primaryKey"`
	RaffleId                 uint      `json:"raffle_id"`
	UserId                   uint      `json:"user_id"`
	ErajayaClubUserId        uint      `json:"erajaya_club_user_id"`
	RaffleProductId          uint      `json:"raffle_product_id"`
	RaffleProductSizeStockId uint      `json:"raffle_product_size_stock_id"`
	QueueNo                  uint      `json:"queue_no"`
	Status                   uint      `json:"status" gorm:"DEFAULT:0"`
	CreatedAt                time.Time `json:"created_at" gorm:"DEFAULT:null"`
	CreatedBy                string    `json:"created_by" gorm:"DEFAULT:null"`
	UpdatedAt                time.Time `json:"updated_at" gorm:"DEFAULT:null"`
	UpdatedBy                string    `json:"updated_by" gorm:"DEFAULT:null"`
	DeletedAt                time.Time `json:"deleted_at" gorm:"DEFAULT:null"`
	DeletedBy                string    `json:"deleted_by" gorm:"DEFAULT:null"`
}
