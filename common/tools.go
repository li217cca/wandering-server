package common

import "math/rand"
import "time"
import "fmt"

var (
	ran *rand.Rand
)

// GetRand ...
func GetRand() *rand.Rand {
	if ran == nil {
		ran = rand.New(rand.NewSource(time.Now().UnixNano()))
		// fmt.Println("GetRand new ran %v", ran)
	}
	return ran
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

// GetTodayLucky ...
func GetTodayLucky() int {
	r := GetRand()
	ans := 0
	for i := 0; i < 10; i++ { // normal day lucky = [0, 100]
		ans += r.Intn(10)
	}
	return ans
}

// Roulette ...
type Roulette []struct {
	Weight int
	Target interface{}
}

// Get ...
func (rou *Roulette) Get() interface{} {
	totWeight := 0
	var tmp []struct {
		WeightPos int
		Target    *interface{}
	}
	for index := range *rou {
		totWeight += (*rou)[index].Weight
		tmp = append(tmp, struct {
			WeightPos int
			Target    *interface{}
		}{
			WeightPos: totWeight,
			Target:    &((*rou)[index].Target),
		})
	}
	r := GetRand()
	dice := r.Int() % totWeight
	for index := range tmp {
		if tmp[index].WeightPos >= dice {
			return *(tmp[index].Target)
		}
	}
	fmt.Printf("Roulette.get error %v", rou)
	return nil
}
