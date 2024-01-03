package hourly

import (
	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInt32Type(t *testing.T) {
	fakeClock := clockwork.NewFakeClock()
	Clock = fakeClock

	var c Int32
	assert.Equal(t, int32(0), c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, int32(1), c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, int32(2), c.Get(), "they should be equal")

	fakeClock.Advance(1 * time.Hour)
	assert.Equal(t, int32(0), c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, int32(1), c.Get(), "they should be equal")

	c.Inc()
	assert.Equal(t, int32(2), c.Get(), "they should be equal")
}
