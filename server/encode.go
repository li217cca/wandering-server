package server

import (
	"encoding/json"
	"wandering-server/model"
)


// 编码消息
func Encode(msg model.Message) ([]byte, error) {
	buf, err := json.Marshal(msg)
	// TODO 添加加密过程
	return buf, err
}

// 解码消息
func Decode(buf []byte) (model.Message, error) {
	msg := model.Message{}
	err := json.Unmarshal(buf, &msg)
	// TODO 添加解密过程
	return msg, err
}
