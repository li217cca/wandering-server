package model

import "testing"
import "time"

func TestParseUserToken(t *testing.T) {
	ID := 123213
	token, _ := NewUserToken(User{ID: ID}, time.Now().Add(time.Millisecond*1000).Unix())
	timer := time.NewTimer(time.Millisecond * 900)
	<-timer.C
	parse, err := ParseUserToken(token)
	if parse != ID || err != nil {
		t.Errorf("\ntoken = %v\nparse = %d\nerr = %v\n", token, parse, err)
	}
}
