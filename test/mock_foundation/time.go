package mock_foundation

import "time"

type MockTime struct {
	Time time.Time
}

const timeLayout = "2006-01-02 15:04:05"

func (m MockTime) Now() time.Time {
	return m.Time
}

func (m MockTime) NowTaipei() time.Time {
	return m.Time
}

func (m MockTime) NowUSEast() time.Time {
	return m.Time
}

func (m MockTime) Sleep(duration time.Duration) {}

// 取得帶入時間
func GetTime(t string) time.Time {
	T, _ := time.Parse(timeLayout, t)
	return T
}
