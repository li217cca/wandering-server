package model

// Item Item struct
type Item struct {
	ID       int     `json:"id"`
	BagID    int     `json:"bag_id"   gorm:"index"`
	Capacity int     `json:"capacity" gorm:"capacity"`
	Weight   float32 `json:"weight"   gorm:"weight"`
}

func (item *Item) delete() error {
	return DB.Where(Item{ID: item.ID}).Delete(&item).Error
}
func (item *Item) commit() error {
	if (DB.Where(Item{ID: item.ID}).RecordNotFound()) {
		return DB.Model(item).Create(&item).Error
	}
	return DB.Where(Item{ID: item.ID}).Update(&item).Error
}

// GetItemByID Get a item by id.
func GetItemByID(ID int) (Item, error) {
	var item Item
	err := DB.Where(Item{ID: ID}).Find(&item).Error
	return item, err
}
