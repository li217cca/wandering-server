package model

// Route Map's edge
type Route struct {
	ID       int
	SourceID int
	TargetID int
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
