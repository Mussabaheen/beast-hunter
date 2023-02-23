package battle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomNew_RangeStats(t *testing.T) {

	stats := NewRandomStats(70, 80, 40, 50, 30, 40, 10, 20, 30, 40)

	assert.GreaterOrEqual(t, stats.Health, 70)
	assert.GreaterOrEqual(t, stats.Strength, 40)
	assert.GreaterOrEqual(t, stats.Defence, 30)
	assert.GreaterOrEqual(t, stats.Speed, 10)
	assert.GreaterOrEqual(t, stats.Luck, 30)
	assert.LessOrEqual(t, stats.Health, 80)
	assert.LessOrEqual(t, stats.Strength, 50)
	assert.LessOrEqual(t, stats.Defence, 40)
	assert.LessOrEqual(t, stats.Speed, 20)
	assert.LessOrEqual(t, stats.Luck, 40)
}
