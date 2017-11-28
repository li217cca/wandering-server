package server

import (
	"sync"
	"time"
	"wandering-server/common"

	"github.com/jinzhu/gorm"

	"wandering-server/model"
)

type gameContainersMap struct {
	sync.RWMutex
	m map[int]*gameContainer
}

var (
	db             *gorm.DB
	gameContainers = gameContainersMap{}
)

// func getTime() *model.Time {
// 	hour := time.Now().Hour()
// 	weather := ""
// 	if hour < 6 {
// 		weather = common.WEATHER_WINTER
// 	} else if hour < 12 {
// 		weather = common.WEATHER_SPRING
// 	} else if hour < 18 {
// 		weather = common.WEATHER_SUMMER
// 	} else {
// 		weather = common.WEATHER_AUTUMN
// 	}
// 	return &model.Time{
// 		int(math.Floor(time.Now().Sub(beginYear.Date).Hours()) / 24),
// 		weather,
// 		hour % 6,
// 	}
// }

func init() {
	db = model.DB
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for {
			<-ticker.C
			gameContainers.RLock()
			for _, ptr := range gameContainers.m {
				ptr.Lucky = common.GetTodayLucky(ptr.Info.ID)
			}
			gameContainers.RUnlock()
		}
	}()
	// if db.Model(model.History{}).Where(model.History{Name: "BEGIN_YEAR"}).Find(&beginYear).RecordNotFound() {
	// 	beginYear = model.History{
	// 		Name: "BEGIN_YEAR",
	// 		Date: time.Date(2017, 8, 16, 0, 0, 0, 0, time.Local),
	// 		//Date: time.Now(),
	// 	}
	// 	db.Create(&beginYear)
	// }

	// go func() {
	// 	ticker := time.NewTicker(time.Second * 1)
	// 	for {
	// 		nowTime = getTime()
	// 		<-ticker.C
	// 	}
	// }()
}
