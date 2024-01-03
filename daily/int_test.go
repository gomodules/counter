package daily

import (
	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntType(t *testing.T) {
	fakeClock := clockwork.NewFakeClock()
	Clock = fakeClock

	var c Int
	assert.Equal(t, 0, c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, 1, c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, 2, c.Get(), "they should be equal")

	fakeClock.Advance(1 * time.Hour)
	assert.Equal(t, 2, c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, 3, c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, 4, c.Get(), "they should be equal")

	fakeClock.Advance(24 * time.Hour)
	assert.Equal(t, 0, c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, 1, c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, 2, c.Get(), "they should be equal")
}
