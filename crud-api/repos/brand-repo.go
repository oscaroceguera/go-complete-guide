package repos

import (
	"fmt"
	"golang-crud-rest-api/entities"
)

type BrandRepo struct {
	brands []entities.Brand
}

func NewBrandRepo() *BrandRepo{
	return &BrandRepo{make([]entities.Brand, 0)}
}

func (b *BrandRepo) Create(partial entities.Brand) entities.Brand{
	newItem := partial
	newItem.ID = uint(len(b.brands)) + 1
	b.brands = append(b.brands, newItem)
	return newItem
}

func (d *BrandRepo) GetList()[]entities.Brand {
	return d.brands
}

func (d *BrandRepo) GetOne(id uint) (entities.Brand, error) {
	for _,it := range d.brands {
 		if it.ID == id {
			return it, nil
		}
 	}

	return entities.Brand{}, fmt.Errorf("key '%d' not found", id)
}

func (d *BrandRepo) Update(id uint, amended entities.Brand) (entities.Brand, error) {
	for i, it := range d.brands {
		if it.ID == id {
			d.brands = append(d.brands[:i], d.brands[i+1:]...)
			d.brands = append(d.brands, amended)
			return amended, nil
		}
	}
	return entities.Brand{}, fmt.Errorf("key '%d' not found", amended.ID)
}

func (b *BrandRepo) DeleteOne(id uint) (bool, error) {
	for i, it := range b.brands {
		if it.ID == id {
			b.brands = append(b.brands[:i], b.brands[i+1:]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("key '%d' not found", id)
}