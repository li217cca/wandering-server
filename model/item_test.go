package model

import (
	"reflect"
	"testing"
)

func TestGetItemByID(t *testing.T) {
	item := Item{
		Name:     "test item",
		Capacity: 12,
	}
	item.commit()
	type args struct {
		ID int
	}
	tests := []struct {
		name     string
		args     args
		wantItem Item
		wantErr  bool
	}{
		{
			name:     "1",
			args:     args{},
			wantItem: Item{},
			wantErr:  true,
		},
		{
			name:     "2",
			args:     args{ID: item.ID},
			wantItem: item,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem, err := GetItemByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItemByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("GetItemByID() = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}
