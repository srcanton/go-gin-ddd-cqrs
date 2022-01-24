package model

// User is the model of users
type Product struct {
	ID   string  `gorm:"primary_key;column:id;type:uuid" json:"id"`
	Name string  `gorm:"size:100;not null;" json:"name"`
	Cost float64 `gorm:"not null;" json:"cost"`
}

func (p *Product) ToArray() map[string]interface{} {
	product := map[string]interface{}{
		"ID":   p.ID,
		"Name": p.Name,
		"Cost": p.Cost,
	}
	return product
}
