package redis

import "time"

type Expiration struct {
	Second uint32
	Minute uint32
	Hour   uint32
	Day    uint32
	Week   uint32
	Month  uint32
}

func NewExpiration() *Expiration {
	return &Expiration{
		Second: 1,
		Minute: 1 * 60,
		Hour:   1 * 60 * 60,
		Day:    1 * 60 * 60 * 24,
		Week:   1 * 60 * 60 * 24 * 7,
		Month:  1 * 60 * 60 * 24 * 31,
	}
}

// 以秒为单位:
func (m *Expiration) Seconds(num uint32) time.Duration {
	return calcSeconds(m.Second, num)
}

// 以分钟为单位:
func (m *Expiration) Minutes(num uint32) time.Duration {
	return calcSeconds(m.Minute, num)
}

// 以小时为单位:
func (m *Expiration) Hours(num uint32) time.Duration {
	return calcSeconds(m.Hour, num)
}

// 以日为单位:
func (m *Expiration) Days(num uint32) time.Duration {
	return calcSeconds(m.Day, num)
}

// 以周为单位:
func (m *Expiration) Weeks(num uint32) time.Duration {
	return calcSeconds(m.Week, num)
}

// 以月为单位:
func (m *Expiration) Months(num uint32) time.Duration {
	return calcSeconds(m.Month, num)
}

//////////////////////////////////////////////////////////////////////////

// 以秒为单位计算:
func calcSeconds(unit uint32, unitNum uint32) time.Duration {
	return time.Second * time.Duration(unit*unitNum)
}
