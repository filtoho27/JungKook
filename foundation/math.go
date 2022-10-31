package foundation

import "math"

// 四捨五入小數第N位
func Round(f float64, n int) float64 {
	shift := math.Pow10(n)
	return math.Floor(f*shift+.5) / shift
}

// 無條件捨去小數第N位
func Floor(f float64, n int) (floorNumber float64) {
	pow := math.Pow10(n)
	floorNumber = math.Floor(Round(f*pow, 4)) / pow
	return
}
