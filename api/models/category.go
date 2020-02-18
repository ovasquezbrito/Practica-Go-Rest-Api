package models

import "errors"

//Category modelo
type Category struct {
	Model
	Description string     `gorm:"size:256; not nul; unique" json:"description"`
	Products    []*Product `gorm:"foreignkey:CategoryID" json:"products"`
}

var (
	ErrCategoryEmptyDescription = errors.New("category.description can't empty")
)

//Validate metodo para que el campo description no este vacío
func (c *Category) Validate() error {
	if c.Description == "" {
		return ErrCategoryEmptyDescription
	}
	return nil
}
