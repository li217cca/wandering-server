package model

import (
	"math"
	"reflect"
	"testing"
	"time"
)

func TestResource_delete(t *testing.T) {
	res := Resource{
		Level: 11,
	}
	res.commit()
	tests := []struct {
		name    string
		res     *Resource
		wantErr bool
	}{
		{
			name:    "1",
			res:     &Resource{},
			wantErr: true,
		},
		{
			name:    "2",
			res:     &res,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.res.delete(); (err != nil) != tt.wantErr {
				t.Errorf("Resource.delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetResourceByID(t *testing.T) {
	res := Resource{
		MapID:   0,
		Type:    1,
		Level:   11,
		PreTime: time.Now(),
	}
	res.commit()
	type args struct {
		ID int
	}
	tests := []struct {
		name    string
		args    args
		wantRes Resource
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{ID: 0},
			wantRes: Resource{},
			wantErr: true,
		},
		{
			name:    "2",
			args:    args{ID: res.ID},
			wantRes: res,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := GetResourceByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("\nGetResourceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotRes.PreTime = tt.wantRes.PreTime
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("\nGetResourceByID() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestResource_recovery(t *testing.T) {
	pre := time.Now()
	pre = pre.Add(-time.Hour * 24)
	tests := []struct {
		name    string
		res     *Resource
		wantRes *Resource
	}{
		{
			name: "1",
			res: &Resource{
				Quantity:      1,
				QuantityLimit: 100,
				RecoverySpeed: 10,
				PreTime:       pre,
			},
			wantRes: &Resource{
				Quantity:      11,
				QuantityLimit: 100,
			},
		},
		{
			name: "2",
			res: &Resource{
				Quantity:      0,
				QuantityLimit: 100,
				RecoverySpeed: 200,
				PreTime:       pre,
			},
			wantRes: &Resource{
				Quantity: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.res.recovery()
			if math.Abs(tt.res.Quantity-tt.wantRes.Quantity) > 1e-6 {
				t.Errorf("\nResource.recovery() = %v, want %v", tt.res, tt.wantRes)
			}
		})
	}
}
