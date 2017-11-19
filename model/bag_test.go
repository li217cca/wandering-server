package model

import (
	"reflect"
	"testing"
)

func TestBag_delete(t *testing.T) {
	bag := NewBag()
	bag2 := NewBag()
	bag2.Items = []Item{
		Item{},
	}
	tests := []struct {
		name    string
		bag     *Bag
		wantErr bool
	}{
		{
			name: "1",
			bag: &Bag{
				CapacityLimit: 12,
			},
			wantErr: true,
		},
		{
			name:    "2",
			bag:     &bag,
			wantErr: false,
		},
		{
			name:    "3",
			bag:     &bag2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bag.delete(); (err != nil) != tt.wantErr {
				t.Errorf("Bag.delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetBagByID(t *testing.T) {
	bag := NewBag()
	bag.Items = []Item{
		Item{
			Name: "test 1",
		},
	}
	bag.commit()
	type args struct {
		ID int
	}
	tests := []struct {
		name    string
		args    args
		wantBag Bag
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{ID: 0},
			wantBag: Bag{},
			wantErr: true,
		},
		{
			name:    "2",
			args:    args{ID: bag.ID},
			wantBag: bag,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBag, err := GetBagByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBagByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBag, tt.wantBag) {
				t.Errorf("GetBagByID() = %v, want %v", gotBag, tt.wantBag)
			}
		})
	}
	bag.delete()
}

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
					Skill{Type: SkillBagCapacityID, Level: 3},
					Skill{Type: SkillBagCapacityID, Level: 9},
					Skill{Type: SkillBagWeightID, Level: 5},
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

func TestBag_commitWithoutChildren(t *testing.T) {
	bag := NewBag()
	bag.CapacityLimit = 100
	tests := []struct {
		name    string
		bag     *Bag
		wantBag Bag
	}{
		{
			name:    "1",
			bag:     &bag,
			wantBag: bag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bag.commitWithoutChildren()
			tmpBag, err := GetBagByID(bag.ID)
			if err != nil {
				t.Errorf("\nBag.commitWithoutChildren Faild %v", err)
			}
			if !reflect.DeepEqual(tmpBag, tt.wantBag) {
				t.Errorf("\nBag.commitWithoutChildren() = %v, want %v", tmpBag, tt.wantBag)
			}
		})
	}
}
