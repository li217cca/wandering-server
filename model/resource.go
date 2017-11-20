package model

import (
	"fmt"
	"math"
	"time"
	"wandering-server/common"
)

// Resource ...
type Resource struct {
	ID    int `json:"id"`
	MapID int `json:"map_id"`
	Type  int `json:"type" gorm:"not null"`

	Level         int `json:"level"`
	Quantity      float64
	QuantityLimit float64
	RecoverySpeed float64   `json:"recovery_speed"` // per day
	PreTime       time.Time `json:"pre_time" gorm:"not null"`
}

/*
Resource const
*/
const (
	// TODO: Add more ResourceType
	ResourceMetalID     = 100
	ResourceRareMetalID = 101
	ResourceWaterID     = 200
	ResourceLegendID    = 700
)

/*
Resource.preSreach
Type: pure
UnitTest: false
*/
func (res *Resource) preSearch() {
	// TODO: Resource.preSearch
}

/*
Resource.recovery
Type: pure
UnitTest: true
*/
func (res *Resource) recovery() {
	if res.RecoverySpeed == 0 {
		res.PreTime = time.Now()
		return
	}
	diff := time.Now().Sub(res.PreTime).Hours() * res.RecoverySpeed / 24.
	res.Quantity += diff
	if res.Quantity > res.QuantityLimit {
		res.Quantity = res.QuantityLimit
	}
	res.PreTime = time.Now()
}

/*
NewResource ...
Type: pure
UnitTest: false
*/
func NewResource(Type int, level int, Quantity float64, QuantityLimit float64, RecoverySpeed float64, PreTime time.Time) Resource {
	res := Resource{
		Type:          Type,
		Level:         level,
		Quantity:      Quantity,
		QuantityLimit: QuantityLimit,
		RecoverySpeed: RecoverySpeed,
		PreTime:       PreTime,
	}
	return res
}

/*
defaultRecoveryDay ...
Type: pure
UnitTest: false
*/
func defaultRecoveryDay(typeID int) float64 {
	switch typeID {
	case ResourceMetalID:
		return 30
	case ResourceWaterID:
		return 1
	case ResourceRareMetalID:
		return 120
	case ResourceLegendID:
		return 365
	}
	fmt.Printf("defaultRecoveryDay Resource typeID = %d not found", typeID)
	return -1
}

/*
bfoQuantity ...
Type: pure
UnitTest: false
*/
func bfoQuantity(typeID int) float64 {
	switch typeID {
	case ResourceMetalID:
		return 1.2
	case ResourceWaterID:
		return 0.8
	case ResourceRareMetalID:
		return 0.3
	case ResourceLegendID:
		return 0.1
	}
	fmt.Printf("defaultRecoveryDay Resource typeID = %d not found", typeID)
	return -1
}

/*
randomResourceType ...
Type: pure
UnitTest: true
*/
func randomResourceType(miracle int, danger int) int {
	r := common.GetRand()
	miracle = miracle/4*3 + r.Intn(miracle/2)
	danger = danger + r.Intn(10) - 5
	if danger < 1 {
		danger = 1
	}
	rou := common.Roulette{
		{
			Weight: 100,
			Target: ResourceMetalID,
		},
		{
			Weight: int(math.Sqrt(float64(miracle))+10*math.Sqrt(float64(danger))) / 9,
			Target: ResourceRareMetalID,
		},
		{
			Weight: 100,
			Target: ResourceWaterID,
		},
		{
			Weight: int(math.Sqrt(float64(miracle)*math.Sqrt(float64(danger)))) / 10.,
			Target: ResourceLegendID,
		},
	}
	return rou.Get().(int)
}

/*
GenerateResource ...
Type: pure
UnitTest: false
*/
func GenerateResource(lucky int, miracle int, danger int) Resource {
	r := common.GetRand()
	typeID := randomResourceType(r.Intn(lucky)*r.Intn(lucky)+miracle, danger)
	quantity := (float64(r.Intn(lucky)) + 51.1) * (float64(r.Intn(lucky)) + 51.1) * bfoQuantity(typeID)
	recDay := defaultRecoveryDay(typeID)
	recSpeed := 0.
	if recDay != -1 {
		recSpeed = quantity / (recDay + math.Sqrt(math.Sqrt(quantity)))
	}
	res := Resource{
		Type:          typeID,
		Level:         danger + r.Intn(10) + r.Intn(10) - 10,
		Quantity:      quantity,
		QuantityLimit: quantity,
		RecoverySpeed: recSpeed,
	}
	return res
}
