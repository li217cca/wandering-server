package model

/*
Bag Bag's struct
*/
type Bag struct {
	ID            int     `json:"id,omitempty"`
	Capacity      int     `json:"capacity,omitempty"`
	CapacityLimit int     `json:"capacity_limit,omitempty"`
	Weight        float64 `json:"weight,omitempty"`
	WeightLimit   float64 `json:"weight_limit,omitempty"`
	Items         []Item  `gorm:"ForeignKey" json:"items,omitempty"`
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
Type: pure
UnitTest: false
*/
func NewBag() Bag {
	bag := Bag{
		Items: []Item{},
	}
	return bag
}
