package model

import (
	"testing"
)

func TestNewMap(t *testing.T) {
	type args struct {
		lucky     float64
		bfoDanger float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{
				lucky:     100,
				bfoDanger: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMap(tt.args.lucky, tt.args.bfoDanger)
			t.Error("\n创世", got.ToString())
			got.Resource.Evolved(5)
			t.Error("\n五分钟", got.ToString())
			got.Resource.Evolved(60)
			t.Error("\n一小时", got.ToString())
			got.Resource.Evolved(1440)
			t.Error("\n一天", got.ToString())
			got.Resource.Evolved(43200)
			t.Error("\n一个月", got.ToString())
		})
	}
}
