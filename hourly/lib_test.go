package hourly

import (
	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInt(t *testing.T) {
	fakeClock := clockwork.NewFakeClock()
	Clock = fakeClock

	var c int
	assert.Equal(t, 0, Get(&c), "they should be equal")

	Inc(&c)
	assert.Equal(t, 1, Get(&c), "they should be equal")

	Inc(&c)
	assert.Equal(t, 2, Get(&c), "they should be equal")

	fakeClock.Advance(1 * time.Hour)
	assert.Equal(t, 0, Get(&c), "they should be equal")

	Inc(&c)
	assert.Equal(t, 1, Get(&c), "they should be equal")

	Inc(&c)
	assert.Equal(t, 2, Get(&c), "they should be equal")
}

func TestInt32(t *testing.T) {
	fakeClock := clockwork.NewFakeClock()
	Clock = fakeClock

	var c int32
	assert.Equal(t, int32(0), Get32(&c), "they should be equal")

	Inc32(&c)
	assert.Equal(t, int32(1), Get32(&c), "they should be equal")

	Inc32(&c)
	assert.Equal(t, int32(2), Get32(&c), "they should be equal")

	fakeClock.Advance(1 * time.Hour)
	assert.Equal(t, int32(0), Get32(&c), "they should be equal")

	Inc32(&c)
	assert.Equal(t, int32(1), Get32(&c), "they should be equal")

	Inc32(&c)
	assert.Equal(t, int32(2), Get32(&c), "they should be equal")
}

func TestInt64(t *testing.T) {
	fakeClock := clockwork.NewFakeClock()
	Clock = fakeClock

	var c int64
	assert.Equal(t, int64(0), Get64(&c), "they should be equal")

	Inc64(&c)
	assert.Equal(t, int64(1), Get64(&c), "they should be equal")

	Inc64(&c)
	assert.Equal(t, int64(2), Get64(&c), "they should be equal")

	fakeClock.Advance(1 * time.Hour)
	assert.Equal(t, int64(0), Get64(&c), "they should be equal")

	Inc64(&c)
	assert.Equal(t, int64(1), Get64(&c), "they should be equal")

	Inc64(&c)
	assert.Equal(t, int64(2), Get64(&c), "they should be equal")
}
