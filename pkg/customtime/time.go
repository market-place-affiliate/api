package customtime

import "time"

var Now = func() time.Time {
	return time.Now()
}
