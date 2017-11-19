package model

import (
	"fmt"
)

// Item Item struct
type Item struct {
	ID    int `gorm:"primary_key"`
	BagID int `gorm:"foreign_key"`

	Name     string  `json:"name" gorm:"not null"`
	Type     int     `json:"type" gorm:"not null"`
	Capacity int     `json:"capacity" gorm:"capacity"`
	Weight   float64 `json:"weight"   gorm:"weight"`
}

// Item const..
const (
	ItemHitPointRegenID = 200
)

/*
Item.delete
Type: not pure
UnitTest: true
*/
func (item *Item) delete() error {
	if num := DB.Where("id = ?", item.ID).Delete(&item).RowsAffected; num != 1 {
		return fmt.Errorf("Item.delete 01\n RowsAffected = %d", num)
	}
	return nil
}

/*
Item.commit
Type: not pure
UnitTest: false
*/
func (item *Item) commit() {
	DB.Where("id = ?", item.ID).FirstOrCreate(&item)
	DB.Model(item).Update(&item)
}

/*
GetItemByID Get a Item{} by Item.ID from Database
Type: not pure
UnitTest: false
*/
func GetItemByID(ID int) (item Item, err error) {
	if err = DB.Model(item).Where("id = ?", ID).Find(&item).Error; err != nil {
		return item, fmt.Errorf("GetItemByID 01\n %v", err)
	}
	return item, nil
}
