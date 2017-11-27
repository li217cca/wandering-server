package model

import (
	"fmt"
	"testing"
	"wandering-server/common"
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
				lucky:     50,
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
			str := ``
			for i := 0; i < 10; i++ {
				lucky := common.FloatF(10, 130)
				pre := got.GenerateQuest(lucky)
				Length := pre.Length
				pre.ID = 1
				str += pre.ToString()
				got.Quests = append(got.Quests, pre)
				for j := 0; j < Length; j++ {
					getDistiny := pre.GetDistiny()
					useDistiny := pre.UseDistiny(getDistiny)
					destinyDiff := getDistiny - useDistiny
					if pre.IsEnd(destinyDiff) {
						str += fmt.Sprintf("\n		[%d%d = %d]: End\n", pre.Destiny, destinyDiff, pre.Destiny+destinyDiff)
						break
					}
					pre = got.GenerateNextQuest(lucky, destinyDiff, &pre)
					pre.ID = j + 2
					str += pre.ToString()
					got.Quests = append(got.Quests, pre)
				}
			}
			t.Errorf("Quests: \n%s", str)
		})
	}
}
