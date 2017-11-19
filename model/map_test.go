package model

import (
	"testing"
	"time"
)

// 测试delete功能
func TestMap_delete(t *testing.T) {
	mp := NewMap(123, 10)
	route := mp.addRoute(100)
	resource := NewResource(mp.ID, 1, 0, 100, 200, 100, time.Now())
	mp.Resources = append(mp.Resources, resource)
	mp.commit()
	tests := []struct {
		name    string
		mp      *Map
		wantErr bool
	}{
		{
			name:    "1",
			mp:      &Map{},
			wantErr: true,
		},
		{
			name: "2",
			mp: &Map{
				ID:        mp.ID,
				Resources: []Resource{Resource{Type: 1}},
				Routes:    []Route{Route{SourceID: mp.ID}},
			},
			wantErr: true,
		},
		{
			name: "3",
			mp: &Map{
				ID:     mp.ID,
				Routes: []Route{Route{SourceID: mp.ID}},
			},
			wantErr: true,
		},
		{
			name:    "4",
			mp:      &mp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idTmp := mp.ID
			if err := tt.mp.delete(); (err != nil) != tt.wantErr {
				t.Errorf("\nMap.delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			tmp, err := GetMapByID(idTmp)
			if !tt.wantErr && err == nil {
				t.Errorf("\nMap.delete() GetMapByID(%d) Map{} = %v existed", idTmp, tmp)
			}
			if !tt.wantErr && !DB.Where("id = ?", route.ID).First(&route).RecordNotFound() {
				t.Errorf("\nTestMap_delete 03 \nRoute{} = %v existed", route)
			}
			if !tt.wantErr && !DB.Where("id = ?", resource.ID).First(&resource).RecordNotFound() {
				t.Errorf("\nTestMap_delete 03 \nResource{} = %v existed", resource)
			}
		})
	}
}

func TestMap_commitWithoutChildren(t *testing.T) {
	tests := []struct {
		name string
		mp   *Map
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mp.commitWithoutChildren()
		})
	}
}

func TestMap_commit(t *testing.T) {
	tests := []struct {
		name string
		mp   *Map
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mp.commit()
		})
	}
}
