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
