package models

type Product struct {
	Id   int64  `gorm:"primaryKey" json:"id"`
	Nama string `gorm:"type:varchar(100)" json:"nama"`
	Desc string `gorm:"type:text(300)" json:"desc"`
}
