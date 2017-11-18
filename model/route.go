package model

import (
	"fmt"
)

// Route Map's edge
type Route struct {
	ID       int `json:"id"`
	SourceID int `json:"source_id"`
	TargetID int `json:"target_id"`
}

func (route *Route) delete() error {
	return DB.Where(Route{ID: route.ID}).Delete(&route).Error
}

// CreateRoute Create map single link to map
func CreateRoute(SourceID int, TargetID int) error {
	if (!DB.Where(Route{SourceID: SourceID, TargetID: TargetID}).RecordNotFound()) {
		return fmt.Errorf("Route %d to %d already existed", SourceID, TargetID)
	}
	return DB.Model(Route{}).Create(&Route{SourceID: SourceID, TargetID: TargetID}).Error
}
