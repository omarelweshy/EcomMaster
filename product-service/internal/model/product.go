package model

type Product struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"size:255;not null" json:"name"`
	Description string     `json:"description,omitempty"`
	Price       float64    `gorm:"not null" json:"price"`
	Stock       int        `gorm:"not null" json:"stock"`
	Categories  []Category `gorm:"many2many:product_categories" json:"categories,omitempty"`
}
