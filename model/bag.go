package model

import (
	"fmt"
)

// Bag 背包结构
type Bag struct {
	ID            int     `json:"id"`
	CapacityLimit int     `json:"capacity_limit"`
	WeightLimit   float32 `json:"weight_limit"`

	Capacity int     `json:"capacity" gorm:"-"`
	Weight   float32 `json:"weight" gorm:"-"`
	Items    []Item  `json:"items" gorm:"-"`
}

// GetBagByID 通过 Bag.ID 从数据库获得 Bag 实体
func GetBagByID(ID int) (bag Bag, err error) {
	if ID == 0 {
		err = fmt.Errorf("No such bag with BagID %d", ID)
		return
	}
	err = DB.Model(Bag{}).Where(Bag{ID: ID}).Find(&bag).Error
	DB.Model(Item{}).Where(Item{BagID: ID}).Find(&bag.Items)
	bag.refreshCapacityWeight()
	return
}

func (bag *Bag) commit() error {
	var err error
	if (DB.Where(Bag{ID: bag.ID}).RecordNotFound()) {
		err = DB.Model(bag).Create(&bag).Error
	} else {
		err = DB.Where(Bag{ID: bag.ID}).Update(&bag).Error
	}
	if err != nil {
		return err
	}
	for index := range bag.Items {
		err = bag.Items[index].commit()
		if err != nil {
			return err
		}
	}
	err = bag.refreshCapacityWeight()
	return err
}

func (bag *Bag) refreshCapacityWeight() error {
	bag.Capacity = 0
	bag.Weight = 0
	for index := range bag.Items {
		bag.Capacity += bag.Items[index].Capacity
		bag.Weight += bag.Items[index].Weight
	}
	return nil
}
