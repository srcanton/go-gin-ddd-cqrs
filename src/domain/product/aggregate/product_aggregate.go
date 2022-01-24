package product_aggregate

import (
	"fmt"
	product_cost "github.com/srcanton/go-gin-ddd-cqrs/src/domain/product/valueobject/cost"
	product_id "github.com/srcanton/go-gin-ddd-cqrs/src/domain/product/valueobject/id"
	product_name "github.com/srcanton/go-gin-ddd-cqrs/src/domain/product/valueobject/name"
	"strconv"
)

// User is an user which has id, email and person info
type Product struct {
	ID   product_id.ProductId     `gorm:"primary_key;column:id;type:uuid" json:"id"`
	Name product_name.ProductName `gorm:"column:name;type:text;size:100;not null;" json:"name"`
	Cost product_cost.ProductCost `gorm:"column:cost;type:float;not null;" json:"cost"`
}

func FromRaw(array map[string]interface{}) (*Product, error) {
	userIdValueObject, _ := product_id.FromValue(fmt.Sprintf("%v", array["ID"]))
	userNameValueObject, _ := product_name.FromValue(fmt.Sprintf("%v", array["Name"]))
	userCost, _ := strconv.ParseFloat(fmt.Sprintf("%v", array["COST"]), 64)
	userCostValueObject, _ := product_cost.FromValue(userCost)
	return &Product{
		ID:   userIdValueObject,
		Name: userNameValueObject,
		Cost: userCostValueObject,
	}, nil
}
