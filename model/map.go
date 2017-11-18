package model

import "wandering-server/common"

// Map ...
type Map struct {
	ID  int    `json:"id"`
	Key string `json:"key"`

	Danger  int `json:"danger"`  // 危险等级
	Miracle int `json:"miracle"` // 奇迹等级
	// Weather  int `json:"weather"`
	Resources []Resource `gorm:"-"`
}

func (mp *Map) delete() error {
	return DB.Where(Map{ID: mp.ID}).Delete(&mp).Error
}
func (mp *Map) commit() error {
	if (DB.Where(Map{ID: mp.ID}).RecordNotFound()) {
		return DB.Model(mp).Create(&mp).Error
	}
	return DB.Where(Map{ID: mp.ID}).Update(&mp).Error
}

// CreateMap 随机生产Map
func CreateMap(lucky int, bfoDanger int) (Map, error) {
	r := common.GetRand()
	mp := Map{
		Key:     common.GenerateKey(8),
		Danger:  bfoDanger + r.Intn(10) - 5,
		Miracle: r.Intn(lucky)*r.Intn(lucky) + r.Intn(lucky),
	}
	mp.commit()
	// TODO: generate resource
	return mp, nil
}

// GetMapByID ...
func GetMapByID(ID int) (mp Map, err error) {
	if err = DB.Where(Map{ID: ID}).Find(&mp).Error; err != nil {
		return
	}
	mp.Resources, err = GetResourcesByMapID(mp.ID)
	return
}
