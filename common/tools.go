package common

import "math/rand"
import "time"
import "fmt"
import "math"

var (
	ran *rand.Rand
)

func init() {
	ran = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GetRand ...
func GetRand() *rand.Rand {
	return ran
}

const (
	randCharList = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

// Sqrt [pure]...
func Sqrt(a int) int {
	return int(math.Sqrt(float64(a)))
}

// Float [pure]
func Float(low int, high int) int {
	if low > high {
		return Float(high, low)
	}
	return ran.Int()%(high-low+1) + low
}

// FloatF [pure]
func FloatF(low float64, high float64) float64 {
	if low > high {
		return ran.Float64()*(low-high) + high
	}
	return ran.Float64()*(high-low) + low
}

// FloatPercent [pure]
func FloatPercent(value int, percent int) int {
	if (ran.Int() & 1) == 0 {
		return value + ran.Int()%(value*percent)/100
	}
	res := value - ran.Int()%(value*percent)/100
	if res < 0 {
		return 0
	}
	return res
}

// FloatPercentF [pure]
func FloatPercentF(value float64, percent float64) float64 {
	diff := ran.Float64() * value * percent * 0.01
	if (ran.Int() & 1) == 0 {
		return value + diff
	}
	res := value - diff
	if res < 0 {
		return 0
	}
	return res
}

// GaussianlRandF10  高斯分布
func GaussianlRandF10(mean float64, diff float64) float64 {
	res := 0.
	for i := 0; i < 10; i++ {
		res += FloatF(mean-diff, mean+diff)
	}
	return res / 10.
}

// BinomialRandF10  二项分布
func BinomialRandF10(mean float64, max float64) float64 {
	res := 0.
	chance := 1 - mean/max
	for i := 0; i < 9; i++ {
		tmp := FloatF(0, max*0.3)
		max -= tmp
		if ran.Float64() >= chance {
			res += tmp
		}
	}
	if ran.Float64() >= chance {
		res += max
	}
	return res
}

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

// RouletteNode ...
type RouletteNode struct {
	Weight int
	Target interface{}
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
