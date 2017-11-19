package model

import (
	"fmt"
)

// Route Map's edge
type Route struct {
	ID       int
	SourceID int
	TargetID int
}

func (route *Route) delete() error {
	if num := DB.Where("id = ?", route.ID).Delete(&route).RowsAffected; num != 1 {
		return fmt.Errorf("\nRoute.delete 01 \nRowsAffected = %d", num)
	}
	return nil
}

func (route *Route) commit() {
	DB.Save(&route)
}

// CreateRoute Create map single link to map
func CreateRoute(SourceID int, TargetID int) error {
	if (!DB.Where(Route{SourceID: SourceID, TargetID: TargetID}).RecordNotFound()) {
		return fmt.Errorf("Route %d to %d already existed", SourceID, TargetID)
	}
	return DB.Model(Route{}).Create(&Route{SourceID: SourceID, TargetID: TargetID}).Error
}

/*
NewRoute ...
Type: not pure
UnitTest: false
*/
func NewRoute(sourceID int, targetID int) Route {
	route := Route{
		SourceID: sourceID,
		TargetID: targetID,
	}
	route.commit()
	return route
}
