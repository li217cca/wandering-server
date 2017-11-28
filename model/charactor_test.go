package model

import (
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
