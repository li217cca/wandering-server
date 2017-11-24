package model

// Route Map's edge
type Route struct {
	ID       int `json:"id,omitempty"`
	SourceID int `json:"source_id,omitempty"`
	TargetID int `json:"target_id,omitempty"`
	Strength int `json:"strength,omitempty"` // 强度
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
	return route
}
