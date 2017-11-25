package model

import "wandering-server/common"
import "time"
import "fmt"
import "math"

// Map ...
type Map struct {
	ID       int       `json:"id,omitempty"`
	Key      string    `json:"key,omitempty" gorm:"not null"`
	Name     string    `json:"name,omitempty" gorm:"not null"`
	Resource Resource  `json:"resources,omitempty"`
	Routes   []Route   `json:"routes,omitempty" gorm:"ForeignKey:source_id"`
	Quests   []Quest   `json:"quests,omitempty"`
	PreTime  time.Time `json:"pre_time,omitempty" gorm:"not null"`
	// Games     []Game     `json:"games,omitempty" gorm:"many2many:game_maps"`
}

/*
NewMap [pure test] Generate Map{} by random, then generate Map.Resources
Type: pure
UnitTest: false
*/
func NewMap(lucky float64, bfoDanger float64) Map { // Normal day & gift lucky = [50, 150]
	r := common.GetRand()
	miracle := math.Pow(lucky, 3)*r.Float64()*r.Float64()/50. + r.Float64()*100
	danger := common.FloatF(bfoDanger-5., bfoDanger+5.)

	mp := Map{
		Key:      common.GenerateKey(8),
		Resource: NewResource(lucky, miracle, danger),
		PreTime:  time.Now(),
	}
	mp.getName(miracle, danger)
	return mp
}

/*
getName [pure]
*/
func (mp *Map) getName(miracle float64, danger float64) {
	mp.Name = common.GetMiraclePrefix(miracle) +
		common.GetDangerPrefix(danger) +
		"大陆"
}

// Search []
func (mp *Map) Search(lucky float64) {
	// TODO: Search.. generate quest
	times := common.Float(3, 6)
	for i := 0; i < times; i++ {
		mp.GenerateQuest(lucky)
	}
}

// ToString [pure]
func (mp *Map) ToString() string {
	str := fmt.Sprintf(`
	Map{ID: %d, Key: "%s", Name: "%s"}%s`,
		mp.ID, mp.Key, mp.Name, mp.Resource.ToString(),
	)
	return str
}
