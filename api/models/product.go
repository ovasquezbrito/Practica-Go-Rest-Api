package models

import "errors"

type ProductStatus uint8

const (
	ProductStatus_Unavailable = 0
	ProductStatus_Available   = 1
)

//Product Model
type Product struct {
	Model
	Name       string        `gorm:"size:512;not null;unique" json:"name"`
	Price      float64       `gorm:"type:decimal(10,2);not null; default:0.0" json:"price"`
	Quantity   uint64        `gorm:"default:0;unsigned" json:"quantity"`
	Status     ProductStatus `gorm:"varcha(1);default:0" json:"status"`
	CategoryID uint64        `gorm:"not null" json:"category_id"`
}

var (
	ErrProductEmptyName = errors.New("product.name can't be empty")
)

//Validate Metodo para validar que el nombre no este vacío
func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrProductEmptyName
	}
	return nil
}
