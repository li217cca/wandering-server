package model

// Party ...
type Party struct {
	ID     int          `json:"id,omitempty"`
	GameID int          `json:"game_id,omitempty"`
	Chars  [6]Charactor `json:"chars,omitempty" gorm:"many2many:party_chars"`
}

// NewPartyFromQuest [pure]
func NewPartyFromQuest(quest *Quest) Party {
	party := Party{}
	return party
}
