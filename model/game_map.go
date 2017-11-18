package model

// GameLinkMap ...
type GameLinkMap struct {
	ID     int `json:"id"`
	GameID int `json:"game_id"`
	MapID  int `json:"map_id"`
}

// GetMapsByGameID ...
func GetMapsByGameID(GameID int) (maps []Map, err error) {
	var links []GameLinkMap
	if err = DB.Where(GameLinkMap{GameID: GameID}).Find(&links).Error; err != nil {
		return maps, err
	}
	var tmp Map
	for _, item := range links {
		tmp, err = GetMapByID(item.MapID)
		if err != nil {
			return maps, err
		}
		maps = append(maps, tmp)
	}
	return maps, err
}
