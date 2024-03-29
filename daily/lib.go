package daily

import (
	"github.com/jonboulle/clockwork"
)

var Clock = clockwork.NewRealClock()

const (
	SuffixLen = 10 // bits.Len(366) + 1
	Suffix    = 0b01111111111
)

func Inc(addr *int) int {
	stored := *addr
	storedTime := int(stored & Suffix)
	curTime := Clock.Now().YearDay()
	if storedTime == curTime {
		*addr = ((stored>>SuffixLen)+1)<<SuffixLen | storedTime
	} else {
		*addr = 1<<SuffixLen | curTime
	}
	return *addr
}

func Get(addr *int) int {
	stored := *addr
	storedTime := int(stored & Suffix)
	curTime := Clock.Now().YearDay()
	if storedTime == curTime {
		return stored >> SuffixLen
	}
	return 0
}

func Inc64(addr *int64) int64 {
	stored := *addr
	storedTime := int(stored & Suffix)
	curTime := Clock.Now().YearDay()
	if storedTime == curTime {
		*addr = ((stored>>SuffixLen)+1)<<SuffixLen | int64(storedTime)
	} else {
		*addr = 1<<SuffixLen | int64(curTime)
	}
	return *addr
}

func Get64(addr *int64) int64 {
	stored := *addr
	storedTime := int(stored & Suffix)
	curTime := Clock.Now().YearDay()
	if storedTime == curTime {
		return stored >> SuffixLen
	}
	return 0
}

func Inc32(addr *int32) int32 {
	stored := *addr
	storedTime := int(stored & Suffix)
	curTime := Clock.Now().YearDay()
	if storedTime == curTime {
		*addr = ((stored>>SuffixLen)+1)<<SuffixLen | int32(storedTime)
	} else {
		*addr = 1<<SuffixLen | int32(curTime)
	}
	return *addr
}

func Get32(addr *int32) int32 {
	stored := *addr
	storedTime := int(stored & Suffix)
	curTime := Clock.Now().YearDay()
	if storedTime == curTime {
		return stored >> SuffixLen
	}
	return 0
}
