package model

import (
	"fmt"
	"time"
)

// Resource ...
type Resource struct {
	ID    int `json:"id"`
	MapID int `json:"map_id"`
	Type  int `json:"type" gorm:"not null"`

	Level         int `json:"level"`
	Quantity      float64
	QuantityLimit float64
	RecoverySpeed float64   `json:"recovery_speed"`
	PreTime       time.Time `json:"pre_time" gorm:"not null"`
}

/*
Resource.commit
Type: not pure
UnitTest: false
*/
func (res *Resource) commit() {
	DB.Save(&res)
}

/*
Resource.delete
Type: not pure
UnitTest: true
*/
func (res *Resource) delete() error {
	if num := DB.Where("id = ?", res.ID).Delete(&res).RowsAffected; num != 1 {
		return fmt.Errorf("\nResource.delete 01 \nRowsAffected = %d", num)
	}
	return nil
}

/*
Resource.recovery
Type: pure
UnitTest: true
*/
func (res *Resource) recovery() {
	diff := time.Now().Sub(res.PreTime).Hours() * res.RecoverySpeed / 24.
	res.Quantity += diff
	if res.Quantity > res.QuantityLimit {
		res.Quantity = res.QuantityLimit
	}
	res.PreTime = time.Now()
}

/*
GetResourceByID Get resource by Resource.ID
Type: not pure
UnitTest: true
*/
func GetResourceByID(ID int) (res Resource, err error) {
	if err = DB.Where("id = ?", ID).First(&res).Error; err != nil {
		return res, fmt.Errorf("\nGetResourceByID 01 \n%v", err)
	}
	return res, nil
}

/*
NewResource ...
Type: not pure
UnitTest: false
*/
func NewResource(mapID int, Type int, level int, Quantity float64, QuantityLimit float64, RecoverySpeed float64, PreTime time.Time) Resource {
	res := Resource{
		MapID:         mapID,
		Type:          Type,
		Level:         level,
		Quantity:      Quantity,
		QuantityLimit: QuantityLimit,
		RecoverySpeed: RecoverySpeed,
		PreTime:       PreTime,
	}
	res.commit()
	return res
}
