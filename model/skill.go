package model

// Skill ...
type Skill struct {
	ID     int `json:"id"`
	CharID int `json:"char_id" gorm:"index"`

	Name  string `json:"name" gorm:"not null;unique"`
	Level int    `json:"level"`
	Exp   int    `json:"exp"`

	ExpLimit int `json:"exp_limit" gorm:"-"`
}

// GetSkillsByCharID ...
func GetSkillsByCharID(CharID int) (skills []Skill, err error) {
	err = DB.Model(Skill{}).Where("char_id = ?", CharID).Find(&skills).Error
	return skills, err
}

func (skill *Skill) commit() error {
	if (DB.Where(Skill{ID: skill.ID}).RecordNotFound()) {
		return DB.Model(skill).Create(&skill).Error
	}
	return DB.Where(Skill{ID: skill.ID}).Update(&skill).Error
}
