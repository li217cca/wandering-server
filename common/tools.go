package common

import "math/rand"
import "time"
import "fmt"
import "math"
import "reflect"

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

// GaussianlRandF3  高斯分布
func GaussianlRandF3(mean float64, diff float64) float64 {
	res := 0.
	for i := 0; i < 3; i++ {
		res += FloatF(mean-diff, mean+diff)
	}
	return res / 3.
}

// BinomialRandF10  二项分布
func _BinomialRandF10(mean float64, max float64) float64 {
	res := 0.
	chance := 1 - mean/max
	for i := 0; i < 99; i++ {
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

// GetTodayLucky [not pure]
func GetTodayLucky(ID int) float64 {
	t := time.Now()
	day := int64(t.Year() * t.YearDay())
	r := rand.New(rand.NewSource(int64(ID) * day))

	return float64(r.Int31()%34 + r.Int31()%34 + r.Int31()%35)
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

// Remove ...
func (rou Roulette) Remove(target interface{}) {
	for index := range rou {
		if reflect.DeepEqual(rou[index].Target, target) {
			rou = append(rou[:index], rou[index+1:]...)
			break
		}
	}
}

// Get ...
func (rou *Roulette) Get() interface{} {
	totWeight := 0
	var tmp []struct {
		WeightPos int
		Target    *interface{}
	}
	for index := range *rou {
		if (*rou)[index].Weight < 1 {
			continue
		}
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
	dice := 0
	if totWeight != 0 {
		dice = r.Int() % totWeight
	}
	for index := range tmp {
		if tmp[index].WeightPos >= dice {
			return *(tmp[index].Target)
		}
	}
	fmt.Printf("Roulette.get error %v", rou)
	if len(*rou) > 0 {
		return (*rou)[0].Target
	}
	return 0
}
