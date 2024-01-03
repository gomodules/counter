package daily

import (
	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInt64Type(t *testing.T) {
	fakeClock := clockwork.NewFakeClock()
	Clock = fakeClock

	var c Int64
	assert.Equal(t, int64(0), c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, int64(1), c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, int64(2), c.Get(), "they should be equal")

	fakeClock.Advance(1 * time.Hour)
	assert.Equal(t, int64(2), c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, int64(3), c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, int64(4), c.Get(), "they should be equal")

	fakeClock.Advance(24 * time.Hour)
	assert.Equal(t, int64(0), c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, int64(1), c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, int64(2), c.Get(), "they should be equal")
}
