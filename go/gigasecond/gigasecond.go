package gigasecond

import (
	"time"
)

const testVersion = 4
const gigasecondDuration = time.Duration(time.Second * 1000000000)

func AddGigasecond(lived time.Time) time.Time {
	return lived.Add(gigasecondDuration)
}
