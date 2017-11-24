package common

/*
GetMiraclePrefix ..
*/
func GetMiraclePrefix(miracle float64) string {
	if miracle < 1000 {
		return ""
	}
	if miracle < 4000 {
		return "罕见的"
	}
	return "奇迹的"
}

/*
GetDangerPrefix ..
*/
func GetDangerPrefix(danger float64) string {
	if danger < 10 {
		return ""
	}
	if danger < 40 {
		return "危险"
	}
	if danger < 100 {
		return "灾难"
	}
	return "灾祸"
}
