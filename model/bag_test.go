package model

import (
	"reflect"
	"testing"
)

func TestBag_calcCapacityWeight(t *testing.T) {
	type args struct {
		skills []Skill
	}
	tests := []struct {
		name    string
		bag     *Bag
		args    args
		wantBag *Bag
	}{
		{
			name: "1",
			bag: &Bag{
				Capacity: 12,
			},
			args: args{
				skills: []Skill{
					Skill{Type: SkillBagCapacityBaseID, Level: 3},
					Skill{Type: SkillBagCapacityBaseID, Level: 9},
					Skill{Type: SkillBagWeightBaseID, Level: 5},
				},
			},
			wantBag: &Bag{
				Capacity:      12,
				CapacityLimit: 12,
				WeightLimit:   5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bag.calcCapacityWeight(tt.args.skills)
			if !reflect.DeepEqual(tt.bag, tt.wantBag) {
				t.Errorf("Bag.calcCapacityWeight(%v) = %v, want %v", tt.args, tt.bag, tt.wantBag)
			}
		})
	}
}
