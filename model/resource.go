package model

import (
	"fmt"
	"math"
	"wandering-server/common"
)

// Resource ...
type Resource struct {
	ID    int `json:"id,omitempty"`
	MapID int `json:"map_id,omitempty"`

	Danger  float64 `json:"danger,omitempty"`
	Miracle float64 `json:"miracle,omitempty"`

	MountResource  float64 `json:"mount_resource,omitempty"`
	PlaneResource  float64 `json:"plane_resource,omitempty"`
	OceanResource  float64 `json:"ocean_resource,omitempty"`
	DebrisResource float64 `json:"debris_resource,omitempty"`

	MetalResource     float64 `json:"metal_resource,omitempty"`
	RareMetalResource float64 `json:"rare_metal_resource,omitempty"`
	FuelResource      float64 `json:"fuel_resource,omitempty"`

	SunnyResource float64 `json:"sunny_resource,omitempty"`
	RainResource  float64 `json:"rain_resource,omitempty"`
	WindResource  float64 `json:"wind_resource,omitempty"`

	PlantResource     float64 `json:"plant_resource,omitempty"`
	PhytozoonResource float64 `json:"phytozoon_resource,omitempty"`
	CarnivoreResource float64 `json:"carnivore_resource,omitempty"`

	CivilizationResource float64 `json:"civilization_resource,omitempty"`

	LegendResource   float64 `json:"legend_resource,omitempty"`   // 传说
	DisasterResource float64 `json:"disaster_resource,omitempty"` // 灾祸
	SanityResource   float64 `json:"sanity_resource,omitempty"`
	MagicResource    float64 `json:"magic_resource,omitempty"`
	ScienceResource  float64 `json:"science_resource,omitempty"`
}

func addInt(a ...int) int {
	res := 0
	for _, i := range a {
		res += i
	}
	if res < 0 {
		return 0
	}
	return res
}

func add(a ...float64) float64 {
	res := 0.
	for _, i := range a {
		res += i
	}
	if res < 0 {
		return 0
	}
	return res
}

// Evolved [pure]
func (res *Resource) Evolved(day int) {
	if day > 36000 {
		day = 36000
	}
	pre := *res
	for i := 0; i < day; i++ {
		MiracleAddition := math.Sqrt(float64(pre.Miracle) / 2000)
		DangerAddition := math.Sqrt(float64(pre.Danger) / 30)
		RareMetalAddition := math.Sqrt(pre.RareMetalResource / 500)
		MetalAddition := math.Sqrt(pre.MetalResource / 2500) // 100 = 0.2; 2500 = 1; 10000 = 2;
		FuelAddition := math.Sqrt(pre.FuelResource / 5000)
		CivilAddition := math.Sqrt(pre.CivilizationResource*0.1) * 0.1 // 1000 = 1; 100000 = 10;
		MountAddition := math.Sqrt(pre.MountResource) * 0.03
		WindAddition := math.Sqrt(pre.WindResource) * 0.03 // 1000 = 1
		SunnyAddition := math.Sqrt(pre.SunnyResource) * 0.03
		RainAddition := math.Sqrt(pre.SunnyResource) * 0.03
		DisasterAddition := math.Sqrt(pre.DisasterResource / 10000)
		res.MountResource = add(
			pre.MountResource,
			pre.MountResource*-0.00001*WindAddition,
			pre.PlaneResource*0.00001,
			pre.MountResource*-0.000001*DisasterAddition,
			pre.DebrisResource*0.000004*CivilAddition,
		)
		res.PlaneResource = add(
			pre.PlaneResource*0.99999,
			pre.MountResource*0.00001*WindAddition,
			pre.OceanResource*0.00001*SunnyAddition,
			pre.PlaneResource*-0.00001*RainAddition,
			pre.PlaneResource*-0.000001*DisasterAddition,
			pre.DebrisResource*0.000003*CivilAddition,
		)
		res.OceanResource = add(
			pre.OceanResource,
			pre.OceanResource*-0.00001*SunnyAddition,
			pre.PlaneResource*0.00001*RainAddition,
			pre.OceanResource*-0.000001*DisasterAddition,
			pre.DebrisResource*0.000003*CivilAddition,
		)
		res.DebrisResource = add(
			pre.DebrisResource,
			pre.MountResource*0.000001*DisasterAddition,
			pre.PlaneResource*0.000001*DisasterAddition,
			pre.OceanResource*0.000001*DisasterAddition,
			pre.DebrisResource*-0.00001*CivilAddition,
		)

		res.MetalResource = add(
			pre.MetalResource*0.999999,
			pre.MountResource*0.00004*MountAddition*MiracleAddition,
			pre.CivilizationResource*-0.000003*CivilAddition,
		)
		res.RareMetalResource = add(
			pre.RareMetalResource*0.999,
			pre.MetalResource*0.00003*MiracleAddition,
			pre.MagicResource*-0.00002*CivilAddition,
		)
		res.FuelResource = add(
			pre.FuelResource*0.999999,
			pre.PlantResource*0.00001*MountAddition,
			pre.CivilizationResource*-0.00001*CivilAddition,
		)

		res.SunnyResource = add(
			pre.SunnyResource*0.998,
			pre.RainResource*0.001,
			pre.WindResource*0.001,
		)
		res.RainResource = add(
			pre.RainResource*0.998,
			pre.SunnyResource*0.001,
			pre.WindResource*0.001,
		)
		res.WindResource = add(
			pre.WindResource*0.998,
			pre.RainResource*0.001,
			pre.SunnyResource*0.001,
		)
		res.SunnyResource = common.FloatPercentF(pre.SunnyResource, 0.2)
		res.RainResource = common.FloatPercentF(pre.RainResource, 0.2)
		res.WindResource = common.FloatPercentF(pre.WindResource, 0.2)

		res.PlantResource = add(
			pre.PlantResource*0.98,
			pre.PlaneResource*0.04*SunnyAddition*RainAddition*WindAddition,
			pre.OceanResource*0.01*SunnyAddition,
			pre.PhytozoonResource*-0.01,
			pre.CivilizationResource*-0.01,
		)
		res.PhytozoonResource = add(
			pre.PhytozoonResource*0.99,
			pre.PlantResource*0.01,
			pre.CarnivoreResource*-0.03,
			pre.CivilizationResource*-0.03,
		)
		res.CarnivoreResource = add(
			pre.CarnivoreResource*0.99,
			pre.PhytozoonResource*0.003,
			pre.CivilizationResource*-0.01,
		)

		CarnivoreAddition := math.Sqrt(pre.CarnivoreResource / 200)
		PhytozoonAddition := math.Sqrt(pre.PhytozoonResource / 500)
		PlantAddition := math.Sqrt(pre.PlantResource / 1000)

		res.CivilizationResource = add(
			pre.CivilizationResource,
			pre.CarnivoreResource*0.000001,
			pre.CarnivoreResource*0.001*CarnivoreAddition*MetalAddition*
				PhytozoonAddition*PlantAddition*CivilAddition,
			pre.CivilizationResource*-0.001*CivilAddition,
			pre.DisasterResource*-0.01*CivilAddition,
		)

		MagicAddition := math.Sqrt(pre.MagicResource / 1000)
		ScienceAddition := math.Sqrt(pre.MagicResource / 1000)

		res.LegendResource = add(
			pre.LegendResource*0.9995,
			pre.CivilizationResource*0.0001*RareMetalAddition,
			pre.ScienceResource*-0.000003*CivilAddition,
			pre.MagicResource*0.000003*CivilAddition,
		)
		res.DisasterResource = add(
			pre.DisasterResource*0.999,
			pre.CivilizationResource*0.00004*CivilAddition*DangerAddition,
			pre.CivilizationResource*-0.0001*MagicAddition,
			pre.CivilizationResource*-0.0001*ScienceAddition,
		)
		res.MagicResource = add(
			pre.MagicResource*0.996,
			pre.OceanResource*0.00001*MiracleAddition,
			pre.DisasterResource*0.00001*MiracleAddition*CivilAddition,
			pre.CivilizationResource*0.0002*RareMetalAddition,
			pre.ScienceResource*-0.0002,
		)
		res.ScienceResource = add(
			pre.ScienceResource*0.999,
			pre.CivilizationResource*0.0002*MetalAddition*FuelAddition,
			pre.MagicResource*-0.0002,
		)
		pre = *res
	}
}

// NewResource [pure test]
func NewResource(lucky float64, miracle float64, danger float64) Resource {
	res := Resource{
		Danger:  danger,
		Miracle: miracle,
	}
	luckyAddition := math.Sqrt(lucky) / 9      // 81
	miracleAddition := math.Sqrt(miracle) / 20 // 3000
	// dangerAddition := math.Sqrt(float64(danger))
	totSize := common.FloatF(200*math.Pow(danger, 0.7), 2000+400*math.Pow(danger, 0.7)*miracleAddition) // [100, X0K]
	res.MountResource = common.FloatF(0, totSize*0.7)
	res.OceanResource = common.FloatF(0, totSize-res.MountResource)
	res.PlaneResource = totSize - res.MountResource - res.OceanResource
	// res.DisasterResource = common.FloatF(0, float64(danger))

	mountAddition := math.Sqrt(res.MountResource) / 70 // 4900
	totSize = common.FloatPercentF(res.MountResource*miracleAddition, 25)
	res.MetalResource = common.FloatF(totSize*0.2, totSize-totSize*0.8)
	res.FuelResource = totSize - res.MetalResource
	res.RareMetalResource = common.FloatPercentF(math.Sqrt(totSize)*miracleAddition*mountAddition*luckyAddition, 30)

	totSize = common.FloatPercentF(math.Sqrt(res.MountResource+res.OceanResource+res.PlaneResource)*50, 25)
	res.SunnyResource = common.FloatF(0, totSize)
	res.WindResource = common.FloatF(0, (totSize - res.SunnyResource))
	res.RainResource = common.FloatF(0, totSize-res.SunnyResource-res.RainResource)

	preDay := 3600 + miracle*10
	if preDay > 360000 {
		preDay = 360000
	}
	preDay = 0
	res.Evolved(preDay)
	return res
}

// ToString [pure]...
func (res *Resource) ToString() string {
	str := fmt.Sprintf(`Resource{ID: %d, MapID: %d, Danger: %.1f, Miracle: %.1f}`,
		res.ID, res.MapID, res.Danger, res.Miracle)
	str += fmt.Sprintf(`
			面积 = %.1f
			山脉 = %.1f  平原 = %.1f  海洋 = %.1f  焦土 = %.1f`,
		res.Area(),
		res.MountResource, res.PlaneResource, res.OceanResource, res.DebrisResource,
	)
	str += fmt.Sprintf(`
			阳光 = %.1f  降雨 = %.1f  风力 = %.1f
			金属 = %.1f  稀金 = %.1f  燃料 = %.1f`,
		res.SunnyResource, res.RainResource, res.WindResource,
		res.MetalResource, res.RareMetalResource, res.FuelResource,
	)
	str += fmt.Sprintf(`
			植物 = %.1f  草食动物 = %.1f  肉食动物 = %.1f
			文明 = %.1f  魔法 = %.1f  科学 = %.1f
			传说 = %.1f  灾祸 = %.1f  SAN = %.1f`,
		res.PlantResource, res.PhytozoonResource, res.CarnivoreResource,
		res.CivilizationResource, res.MagicResource, res.ScienceResource,
		res.LegendResource, res.DisasterResource, res.SanityResource,
	)
	return str
}

// Area [pure]
func (res *Resource) Area() float64 {
	return res.MountResource + res.PlaneResource + res.OceanResource + res.DebrisResource
}
