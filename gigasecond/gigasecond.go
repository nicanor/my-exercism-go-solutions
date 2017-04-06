// Package gigaseconds make awesome calculus.
package gigasecond

import "time"

const testVersion = 4

// AddGigasecond adds a gigasecond to given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000000000000)
}
