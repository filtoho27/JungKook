package foundation

import (
	"time"
)

type RealTime struct{}

var rt *RealTime = &RealTime{}

func GetRealTime() *RealTime {
	if rt == nil {
		rt = &RealTime{}
	}
	return rt
}

func (t RealTime) Now() time.Time {
	return time.Now()
}

func (t RealTime) NowTaipei() time.Time {
	return time.Now().UTC().Add(time.Hour * +8)
}

func (t RealTime) NowUSEast() time.Time {
	return time.Now().UTC().Add(time.Hour * -4)
}

func (t RealTime) Sleep(duration time.Duration) {
	time.Sleep(duration)
}
