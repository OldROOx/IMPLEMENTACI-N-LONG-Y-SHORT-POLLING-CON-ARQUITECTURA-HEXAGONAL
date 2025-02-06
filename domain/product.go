package domain

type Product struct {
	ID     uint    `gorm:"primaryKey"`
	Name   string  `gorm:"size:100"`
	Price  float64 `gorm:"type:decimal(10,2)"`
	UserID uint
}
