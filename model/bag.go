package model

type Bag struct {
	ID            int     `json:"id"`
	Capacity      int     `json:"capacity";gorm:"-"`
	CapacityLimit int     `json:"capacity_limit"`
	WeightLimit   float32 `json:"weight_limit"`
	Weight        float32 `json:"weight";gorm:"-"`
	Items         []Item  `json:"items";gorm:"-"`
}

func GetBagByID(ID int) (bag Bag) {
	if ID == 0 {
		return bag
	}
	DB.Model(Bag{}).Where(Bag{ID: ID}).Find(&bag)
	DB.Model(Item{}).Where(Item{BagID: ID}).Find(&bag.Items)
	return bag
}

func (bag *Bag) Commit() {
	DB.Model(Bag{}).Where(Bag{ID: bag.ID}).Update(&bag)
}
func (bag *Bag) CommitWithItems() {
	bag.Commit()
	for index := range bag.Items {
		bag.Items[index].Commit()
	}
}