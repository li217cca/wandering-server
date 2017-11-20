package model

import (
	"math"
	"testing"
	"time"
)

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

func Test_randomResourceType(t *testing.T) {
	type args struct {
		miracle int
		danger  int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{miracle: 7000, danger: 1},
		},
		{
			name: "1",
			args: args{miracle: 7000, danger: 50},
		},
		{
			name: "1",
			args: args{miracle: 7000, danger: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			return
			set := map[int]int{}
			for i := 0; i < 100000; i++ {
				set[randomResourceType(tt.args.miracle, tt.args.danger)]++
			}
			t.Errorf("\nMiracle %d Danger %d = Set %v", tt.args.miracle, tt.args.danger, set)
		})
	}
}
