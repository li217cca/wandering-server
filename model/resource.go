package model

import (
	"time"
)

// Resource ...
type Resource struct {
	ID    int `json:"id"`
	MapID int `json:"map_id"`
	Type  int `json:"type"`

	Level         int       `json:"level"`
	Quantity      float64   `json:"quantity"`
	QuantityLimit float64   `json:"quantity_limit"`
	RecoverySpeed float64   `json:"recovery_speed"`
	PreTime       time.Time `json:"pre_time"`
}

func (res *Resource) commit() error {
	if (DB.Where(Resource{ID: res.ID}).RecordNotFound()) {
		return DB.Model(res).Create(&res).Error
	}
	return DB.Where(Resource{ID: res.ID}).Update(&res).Error
}
func (res *Resource) delete() error {
	return DB.Where(Resource{ID: res.ID}).Delete(&res).Error
}

// 恢复资源
func (res *Resource) recovery() error {
	res.Quantity += time.Now().Sub(res.PreTime).Hours() * res.RecoverySpeed / 24.
	res.PreTime = time.Now()
	if res.Quantity > res.QuantityLimit {
		res.Quantity = res.QuantityLimit
	}
	return res.commit()
}

// GetResourcesByMapID 获得所有关于MapID的资源，并自动恢复资源
func GetResourcesByMapID(MapID int) (ress []Resource, err error) {
	err = DB.Where(Resource{MapID: MapID}).Find(&ress).Error
	for index := range ress {
		if err = ress[index].recovery(); err != nil {
			return
		}
	}
	return
}
