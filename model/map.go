package model

import "wandering-server/common"
import "math"

// Map ...
type Map struct {
	ID        int        `json:"id,omitempty"`
	Key       string     `json:"key,omitempty" gorm:"not null"`
	Danger    int        `json:"danger,omitempty"`  // 危险等级
	Miracle   int        `json:"miracle,omitempty"` // 奇迹等级
	Resources []Resource `json:"resources,omitempty"`
	Routes    []Route    `json:"routes,omitempty" gorm:"ForeignKey:source_id"`
	Quests    []Quest    `json:"quests,omitempty"`
}

/*
NewMap Generate Map{} by random, then generate Map.Resources
Type: pure
UnitTest: false
*/
func NewMap(lucky int, bfoDanger int) Map { // Normal day & gift lucky = [50, 150]
	r := common.GetRand()
	mp := Map{
		Key:     common.GenerateKey(8),
		Danger:  bfoDanger + r.Intn(10) - 5,
		Miracle: r.Intn(lucky)*r.Intn(lucky)*r.Intn(lucky)/50 + r.Intn(lucky)/2,
	}
	if mp.Danger < 1 {
		mp.Danger = 1
	}
	mp.generateResources(r.Intn(int(math.Sqrt(float64(lucky)))), lucky)
	return mp
}

/*
Map.generateResources
Type: pure
UnitTest: false
*/
func (mp *Map) generateResources(number int, lucky int) {
	for i := 0; i < number; i++ {
		// TODO: generateResources
		tmp := GenerateResource(lucky, mp.Miracle, mp.Danger)
		flag := false
		for index := range mp.Resources {
			if mp.Resources[index].Type == tmp.Type {
				mp.Resources[index].Level = (mp.Resources[index].Level + tmp.Level) / 2
				mp.Resources[index].Quantity += tmp.Quantity
				mp.Resources[index].QuantityLimit += tmp.QuantityLimit
				mp.Resources[index].RecoverySpeed += tmp.RecoverySpeed
				flag = true
				break
			}
		}
		if !flag {
			mp.Resources = append(mp.Resources, tmp)
		} else {
			//
		}
	}
}

/*
Map.research
Type: pure
UnitTest: false
*/
func (mp *Map) research(lucky int) {
	// TODO: research.. generate quest
	r := common.GetRand()
	times := r.Intn(4) + 2
	for i := 0; i < times; i++ {
		mp.generateQuest(lucky)
	}
}
