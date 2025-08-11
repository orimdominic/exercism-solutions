package clock

import (
	"fmt"
	"math"
)

// Define the Clock type here.
type Clock struct {
	h, m int
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func New(h, m int) Clock {
	mins := (h * 60) + m

	if mins < 0 {
		factor := math.Ceil(float64(mins*-1) / float64(24*60))
		mins = (int(factor) * 24 * 60) + mins
		mMod := mins % 60
		hMod := int(mins/60) % 24

		return Clock{
			m: mMod,
			h: hMod,
		}
	}

	mMod := mins % 60
	hMod := int(mins/60) % 24

	return Clock{
		m: mMod,
		h: hMod,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	m = ((c.h * 60) + c.m) - m
	return New(0, m)
}
