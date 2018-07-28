package clock

import (
	"fmt"
)

type clock struct {
	hour int
	minute int	
}

func New(hour int, minute int) clock {
    var h_temp int = 0
	if (minute / 60) != 0 {
		h_temp = minute / 60
	}
	if (minute % 60) <= 0 && minute < 0 && minute != -60 {
		h_temp += -1
	}

	h := (((hour + h_temp) % 24) + 24) % 24
	m := ((minute % 60) + 60) % 60 
	
	return clock{h, m}
}

func (c clock) String() string {
	var h_format string
	var m_format string

	var h_temp int = 0

    if (c.minute / 60) != 0 {
		h_temp = c.minute / 60
	}

	h := (((c.hour + h_temp) % 24) + 24) % 24
	m := ((c.minute % 60) + 60) % 60
	if h > 9 {
		h_format = fmt.Sprintf("%d", h)
	} else {
		h_format = fmt.Sprintf("0%d", h)
	}

	if m > 9 {
		m_format = fmt.Sprintf("%d", m)
	} else {
		m_format = fmt.Sprintf("0%d", m)
	}

	return fmt.Sprintf("%s:%s", h_format, m_format) 
 
}

func (c1 clock) Equal(c2 clock) bool {
	return c1.String() == c2.String()
}

func (c clock) Add(minute int) clock {
	return New(c.hour, c.minute + minute)
}

func (c clock) Subtract(minute int) clock {
	return New(c.hour, c.minute - minute)
}

func main() {
	clock := New(10, 30)
	clock = clock.Subtract(90)
	fmt.Println(clock.String())
}