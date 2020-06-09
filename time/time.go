package time

import (
	"context"
	"database/sql/driver"
	"strconv"
	stdTime "time"
)

////////////////////////////////////////////////////////////////////////////////////////

// Time be used to MySql timestamp converting.
type Time int64

// Scan scan time.
func (m *Time) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case stdTime.Time:
		*m = Time(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		*m = Time(i)
	}
	return
}

// Value get time value.
func (m Time) Value() (driver.Value, error) {
	return stdTime.Unix(int64(m), 0), nil
}

// Time get time.
func (m Time) Time() stdTime.Time {
	return stdTime.Unix(int64(m), 0)
}

////////////////////////////////////////////////////////////////////////////////////////

// Duration be used toml unmarshal string time, like 1s, 500ms.
// extend for toml parse.
type Duration stdTime.Duration

// UnmarshalText unmarshal text to duration.
func (m *Duration) UnmarshalText(text []byte) error {
	tmp, err := stdTime.ParseDuration(string(text))
	if err == nil {
		*m = Duration(tmp)
	}
	return err
}

// Shrink will decrease the duration by comparing with context's timeout duration
// and return new timeout\context\CancelFunc.
func (m Duration) Shrink(c context.Context) (Duration, context.Context, context.CancelFunc) {
	if deadline, ok := c.Deadline(); ok {
		if expire := stdTime.Until(deadline); expire < stdTime.Duration(m) {
			// deliver small timeout
			return Duration(expire), c, func() {}
		}
	}
	ctx, cancel := context.WithTimeout(c, stdTime.Duration(m))
	return m, ctx, cancel
}
