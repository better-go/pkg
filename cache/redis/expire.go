package redis

import "time"

type Expiration struct {
	second uint32
	minute uint32
	hour   uint32
	day    uint32
	week   uint32
	month  uint32
}

func NewExpiration() *Expiration {
	return &Expiration{
		second: 1,
		minute: 1 * 60,
		hour:   1 * 60 * 60,
		day:    1 * 60 * 60 * 24,
		week:   1 * 60 * 60 * 24 * 7,
		month:  1 * 60 * 60 * 24 * 31,
	}
}

// 以秒为单位:
func (m *Expiration) Seconds(num uint32) time.Duration {
	return calcSeconds(m.second, num)
}

// 以分钟为单位:
func (m *Expiration) Minutes(num uint32) time.Duration {
	return calcSeconds(m.minute, num)
}

// 以小时为单位:
func (m *Expiration) Hours(num uint32) time.Duration {
	return calcSeconds(m.hour, num)
}

// 以日为单位:
func (m *Expiration) Days(num uint32) time.Duration {
	return calcSeconds(m.day, num)
}

// 以周为单位:
func (m *Expiration) Weeks(num uint32) time.Duration {
	return calcSeconds(m.week, num)
}

// 以月为单位:
func (m *Expiration) Months(num uint32) time.Duration {
	return calcSeconds(m.month, num)
}

//////////////////////////////////////////////////////////////////////////

// 以秒为单位计算:
func calcSeconds(unit uint32, unitNum uint32) time.Duration {
	return time.Second * time.Duration(unit*unitNum)
}
