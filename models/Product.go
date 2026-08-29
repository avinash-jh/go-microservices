package models

// Product defines the model and schema specifications for GORM.
type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"type:varchar(100);not null"`
	Description string  `json:"description" gorm:"type:varchar(500)"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null"`
}

// TableName tells GORM the exact name of the table to create inside database 'db1'.
func (Product) TableName() string {
	return "productdb"
}
