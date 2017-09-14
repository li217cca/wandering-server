package model

type Item struct {
	ID    int `json:"id"`
	BagID int `json:"bag_id";gorm:"index"`

	Number int `json:"number"`
}

func (i *Item) Commit() {
	DB.Model(i).Where(Item{ID: i.ID}).Update(&i)
}