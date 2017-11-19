package model

import "testing"

func TestSkill_delete(t *testing.T) {
	skill := NewSkill("body", SkillHitPointID, 1)
	tests := []struct {
		name    string
		skill   *Skill
		wantErr bool
	}{
		{
			name:    "1",
			skill:   &Skill{},
			wantErr: true,
		},
		{
			name:    "2",
			skill:   &skill,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.skill.delete(); (err != nil) != tt.wantErr {
				t.Errorf("Skill.delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
