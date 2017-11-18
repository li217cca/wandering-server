package common

import "math/rand"
import "time"

// GetRand ...
func GetRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

const (
	randCharList = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

// GenerateKey ...
func GenerateKey(len int) string {
	str := []byte{}
	ran := GetRand()
	for i := 0; i < len; i++ {
		str = append(str, randCharList[ran.Int31()%36])
	}
	return string(str)
}
