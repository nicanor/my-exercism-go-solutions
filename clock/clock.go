package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	h += m / 60
	m = m % 60

	if m < 0 {
		m += 60
		h = h - 1
	}

	if h < 0 {
		h = 24 + h%24
	}

	h = h % 24
	m = m % 60

	return Clock{h, m}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}
