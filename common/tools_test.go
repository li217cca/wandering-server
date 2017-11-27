package common

import (
	"testing"
)

func TestGetTodayLucky(t *testing.T) {
	type args struct {
		ID int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"1",
			args{
				123,
			},
			65,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTodayLucky(tt.args.ID); got != tt.want {
				t.Errorf("GetTodayLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}
