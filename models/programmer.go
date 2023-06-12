package models

type Programmer struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(300)" json:"name"`
	Keahlian string `gorm:"type:text" json:"keahlian"`
}
