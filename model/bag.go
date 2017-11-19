package model

import (
	"fmt"
)

/*
Bag Bag's struct
*/
type Bag struct {
	ID            int `json:"id"`
	Capacity      int
	CapacityLimit int
	Weight        float64
	WeightLimit   float64
	Items         []Item
}

/*
GetBagByID Get a Bag{} by Bag.ID from Database
Type: not pure
UnitTest: true
*/
func GetBagByID(ID int) (bag Bag, err error) {
	if err = DB.Where("id = ?", ID).Find(&bag).Error; err != nil {
		return bag, fmt.Errorf("GetBagByID 01\n %v", err)
	}
	DB.Model(bag).Related(&bag.Items)
	bag.refreshCapacityWeight()
	return bag, nil
}

/*
Bag.commitWithoutChildren Commit Bag{} to Database
Type: not pure
UnitTest: true
*/
func (bag *Bag) commitWithoutChildren() {
	DB.Model(bag).Save(&bag)
}

/*
Bag.commit Commit Bag{} & Bag.Items to Database
Type: not pure
UnitTest: false
*/
func (bag *Bag) commit() {
	for index := range bag.Items {
		bag.Items[index].commit()
	}
	bag.refreshCapacityWeight()
	bag.commitWithoutChildren()
}

/*
Bag.delete Delete Bag{} and Bag.Items, commit to Database
Type: not pure
UnitTest: true
*/
func (bag *Bag) delete() error {
	for index := range bag.Items {
		if err := bag.Items[index].delete(); err != nil {
			return fmt.Errorf("Bag.delete 01\n %v", err)
		}
	}
	if num := DB.Model(bag).Where("id = ?", bag.ID).Delete(&bag).RowsAffected; num != 1 {
		return fmt.Errorf("Bag.delete 02\n RowsAffected = %d", num)
	}
	return nil
}

/*
refreshCapacityWeight
Type: pure
UnitTest: false
*/
func (bag *Bag) refreshCapacityWeight() {
	bag.Capacity = 0
	bag.Weight = 0
	for index := range bag.Items {
		bag.Capacity += bag.Items[index].Capacity
		bag.Weight += bag.Items[index].Weight
	}
}

/*
calcCapacityWeight
Type: pure
UnitTest: true
*/
func (bag *Bag) calcCapacityWeight(skills []Skill) {
	bag.CapacityLimit = 0
	bag.WeightLimit = 0
	for index := range skills {
		bag.CapacityLimit += skills[index].preCalcBagCapacity()
		bag.WeightLimit += skills[index].preCalcBagWeight()
	}
}

/*
NewBag New a Bag{} & commit to Database
Type: not pure
UnitTest: false
*/
func NewBag() Bag {
	bag := Bag{
		Items: []Item{},
	}
	bag.commit()
	return bag
}
