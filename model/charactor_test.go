package model

import (
	"reflect"
	"testing"
)

func Test_limitAddInt(t *testing.T) {
	type args struct {
		value *int
		limit *int
		inc   int
	}
	value := 50
	limit := 100
	limit0 := 0
	tests := []struct {
		name     string
		args     args
		wantDiff int
		wantErr  bool
	}{
		{
			"1",
			args{
				&value,
				&limit,
				260,
			},
			3,
			false,
		},
		{
			"2",
			args{
				&value,
				&limit,
				-260,
			},
			-3,
			false,
		},
		{
			"3",
			args{
				&value,
				&limit0,
				10,
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDiff, err := limitAddInt(tt.args.value, tt.args.limit, tt.args.inc)
			if (err != nil) != tt.wantErr {
				t.Errorf("limitAddInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDiff != tt.wantDiff {
				t.Errorf("limitAddInt() = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}

func Test_limitAddFloat64(t *testing.T) {
	type args struct {
		value *float64
		limit *float64
		inc   float64
	}
	value := 50.
	limit := 100.
	limit0 := 0.
	tests := []struct {
		name     string
		args     args
		wantDiff int
		wantErr  bool
	}{
		{
			"1",
			args{
				&value,
				&limit,
				260.1,
			},
			3,
			false,
		},
		{
			"2",
			args{
				&value,
				&limit,
				-260,
			},
			-3,
			false,
		},
		{
			"3",
			args{
				&value,
				&limit0,
				10,
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDiff, err := limitAddFloat64(tt.args.value, tt.args.limit, tt.args.inc)
			if (err != nil) != tt.wantErr {
				t.Errorf("limitAddFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDiff != tt.wantDiff {
				t.Errorf("limitAddFloat64() = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}

func TestCharactor_refreshHitPoint(t *testing.T) {
	type args struct {
		Skills []Skill
	}
	tests := []struct {
		name     string
		char     *Charactor
		args     args
		wantChar *Charactor
	}{
		{
			name: "1",
			char: &Charactor{
				HitPoint:      50,
				HitPointLimit: 100,
			},
			args: args{
				Skills: []Skill{
					Skill{
						Type:  SkillHitPointID,
						Level: 10,
					},
				},
			},
			wantChar: &Charactor{
				HitPoint:      55,
				HitPointLimit: 110,
			},
		},
		{
			name: "2",
			char: &Charactor{
				HitPoint:      10,
				HitPointLimit: 100,
			},
			args: args{
				Skills: []Skill{
					Skill{
						Type:  SkillHitPointID,
						Level: 0,
					},
				},
			},
			wantChar: &Charactor{
				HitPoint:      1,
				HitPointLimit: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.char.refreshHitPoint(tt.args.Skills)
			if !reflect.DeepEqual(tt.char, tt.wantChar) {
				t.Errorf("Bag.refreshHitPoint(%v) = %v, want %v", tt.args, tt.char, tt.wantChar)
			}
		})
	}
}

func TestCharactor_delete(t *testing.T) {
	char := NewCharactor(0, CharNotInTeam, []Skill{
		Skill{
			Name: "test skill 01",
		},
	})
	char2 := NewCharactor(0, CharNotInTeam, []Skill{})
	char2.Skills = []Skill{Skill{Name: "Test Bug"}}
	tests := []struct {
		name    string
		char    *Charactor
		wantErr bool
	}{
		{
			name:    "1",
			char:    &Charactor{},
			wantErr: true,
		},
		{
			name:    "2",
			char:    &char,
			wantErr: false,
		},
		{
			name:    "3",
			char:    &char2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.char.delete(); (err != nil) != tt.wantErr {
				t.Errorf("Charactor.delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
