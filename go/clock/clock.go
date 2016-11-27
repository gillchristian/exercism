package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	minutes, rest := getRolledOverTime(minute, 60)
	hours, _ := getRolledOverTime(hour+rest, 24)
	return Clock{hours, minutes}
}

func (c Clock) String() string {
	return fmt.Sprintf(
		"%v:%v", timeToSring(c.hour),
		timeToSring(c.minute),
	)
}

func (c Clock) Add(minutes int) Clock {
	var rest int
	c.minute, rest = getRolledOverTime(minutes+c.minute, 60)
	c.hour, _ = getRolledOverTime(c.hour+rest, 24)
	return c
}

func timeToSring(val int) string {
	if val < 10 {
		return fmt.Sprintf("0%v", val)
	}
	return fmt.Sprintf("%v", val)
}

func getRolledOverTime(val, cap int) (rest, rolledOverTimes int) {
	if val < cap && val >= 0 {
		return val, 0
	}
	rest = val % cap
	rolledOverTimes = (val - rest) / cap

	if val >= 0 {
		return rest, rolledOverTimes
	}
	if rest != 0 {
		rest = rest + cap
	}
	return rest, rolledOverTimes - 1
}
