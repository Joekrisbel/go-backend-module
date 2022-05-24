package models

import (
	"time"

	"gorm.io/gorm"
)

type Raffle struct {
	Id                    uint      `json:"id;primaryKey"`
	Name                  string    `json:"name"`
	StartDateRegistration time.Time `json:"start_date_registration"`
	EndDateRegistration   time.Time `json:"end_date_registration"`
	AnnouncementDate      time.Time `json:"announcement_date"`
	StartDatePay          time.Time `json:"start_date_pay"`
	EndDatePay            time.Time `json:"end_date_pay"`
	Banner                string    `json:"banner"`
	BannerMobile          string    `json:"banner_mobile"`
	Copyright             string    `json:"copyright" gorm:"type:text"`
	Slug                  string    `json:"slug"`
	Status                uint      `json:"status" gorm:"DEFAULT:0"`
	CreatedAt             time.Time `json:"created_at" gorm:"DEFAULT:null"`
	CreatedBy             string    `json:"created_by" gorm:"DEFAULT:null"`
	UpdatedAt             time.Time `json:"updated_at" gorm:"DEFAULT:null"`
	UpdatedBy             string    `json:"updated_by" gorm:"DEFAULT:null"`
	DeletedAt             time.Time `json:"deleted_at" gorm:"DEFAULT:null"`
	DeletedBy             string    `json:"deleted_by" gorm:"DEFAULT:null"`
}

func (raffle *Raffle) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Raffle{}).Count(&total)

	return total
}

func (raffle *Raffle) Take(db *gorm.DB, limit int, offset int) interface{} {
	var raffles []Raffle

	db.Offset(offset).Limit(limit).Find(&raffles)

	return raffles
}

type Raffle_New struct {
	Id                    uint      `json:"id;primaryKey"`
	Name                  string    `json:"name"`
	StartDateRegistration time.Time `json:"start_date_registration"`
	EndDateRegistration   time.Time `json:"end_date_registration"`
	AnnouncementDate      time.Time `json:"announcement_date"`
	StartDatePay          time.Time `json:"start_date_pay"`
	EndDatePay            time.Time `json:"end_date_pay"`
	Banner                string    `json:"banner"`
	Copyright             string    `json:"copyright"`
	Slug                  string    `json:"slug"`
	CreatedAt             time.Time `json:"created_at" gorm:"DEFAULT:null"`
	CreatedBy             string    `json:"created_by" gorm:"DEFAULT:null"`
	UpdatedAt             time.Time `json:"updated_at" gorm:"DEFAULT:null"`
	UpdatedBy             string    `json:"updated_by" gorm:"DEFAULT:null"`
	DeletedAt             time.Time `json:"deleted_at" gorm:"DEFAULT:null"`
	DeletedBy             string    `json:"deleted_by" gorm:"DEFAULT:null"`
}
