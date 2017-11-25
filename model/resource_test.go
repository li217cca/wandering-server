package model

import (
	"testing"
)

func TestNewResource(t *testing.T) {
	type args struct {
		lucky   float64
		miracle float64
		danger  float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				lucky:   70,
				miracle: 5000,
				danger:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get := NewResource(tt.args.lucky, tt.args.miracle, tt.args.danger)
			for i := 0; i < 5; i++ {
				get.Evolved(3600)
				if get.CivilizationResource < 1e-3 ||
					get.CarnivoreResource > get.PhytozoonResource ||
					get.PhytozoonResource > get.PlantResource {
					t.Error("\nResource.Evolved Error\n", get.ToString())
				}
			}
		})
	}
}
