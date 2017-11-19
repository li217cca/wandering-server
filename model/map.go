package model

import "wandering-server/common"
import "fmt"

// Map ...
type Map struct {
	ID  int    `json:"id"`
	Key string `json:"key"`

	Danger  int `json:"danger"`  // 危险等级
	Miracle int `json:"miracle"` // 奇迹等级
	// Weather  int `json:"weather"`
	Resources []Resource
	Routes    []Route `gorm:"ForeignKey:source_id"`
}

/*
Map.delete
Type: not pure
UnitTest: true
*/
func (mp *Map) delete() error {
	for index := range mp.Resources {
		if err := mp.Resources[index].delete(); err != nil {
			return fmt.Errorf("\nMap.delete 01 %v", err)
		}
	}
	for index := range mp.Routes {
		if err := mp.Routes[index].delete(); err != nil {
			return fmt.Errorf("\nMap.delete 01 %v", err)
		}
	}
	if num := DB.Where("id = ?", mp.ID).Delete(&mp).RowsAffected; num != 1 {
		return fmt.Errorf("\nMap.delete 01\n RowsAffected = %d", num)
	}
	return nil
}
func (mp *Map) commitWithoutChildren() {
	DB.Save(&mp)
}

/*
Map.commit ...
Type: not pure
UnitTest: true
*/
func (mp *Map) commit() {
	for index := range mp.Resources {
		mp.Resources[index].commit()
	}
	for index := range mp.Routes {
		mp.Routes[index].commit()
	}
	mp.commitWithoutChildren()
}

func (mp *Map) addRoute(targetID int) Route {
	if mp.ID == 0 {
		mp.commit()
	}
	route := NewRoute(mp.ID, targetID)
	mp.Routes = append(mp.Routes, route)
	return route
}

// NewMap 随机生产Map
func NewMap(lucky int, bfoDanger int) Map {
	r := common.GetRand()
	mp := Map{
		Key:     common.GenerateKey(8),
		Danger:  bfoDanger + r.Intn(10) - 5,
		Miracle: r.Intn(lucky)*r.Intn(lucky)/4 + r.Intn(lucky),
	}
	mp.commit()
	// TODO: generate resource
	return mp
}

// GetMapByID ...
func GetMapByID(ID int) (mp Map, err error) {
	if err = DB.Where(Map{ID: ID}).Find(&mp).Error; err != nil {
		return mp, fmt.Errorf("\nGetMapByID 01 \n%v", err)
	}
	DB.Model(&mp).Related(&mp.Resources)

	return mp, nil
}
