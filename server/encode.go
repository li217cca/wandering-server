package server

import (
	"encoding/json"
	"wandering-server/model"
)

func Encode(msg model.Message) ([]byte, error) {
	buf, err := json.Marshal(msg)
	return buf, err
}

func Decode(buf []byte) (model.Message, error) {
	msg := model.Message{}
	err := json.Unmarshal(buf, &msg)
	return msg, err
}
