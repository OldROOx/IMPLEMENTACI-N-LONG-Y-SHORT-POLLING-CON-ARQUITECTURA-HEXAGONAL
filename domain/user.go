package domain

type User struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `gorm:"size:100"`
	Email    string    `gorm:"unique;size:100"`
	Products []Product `gorm:"foreignKey:UserID"`
}
