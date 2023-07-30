package services

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"yuka-case/pkg/services"
)

func TestDistance(t *testing.T) {
	r := services.Routing{}

	lat1 := 40.748817
	lon1 := -73.985428
	lat2 := 40.7128
	lon2 := -74.0060
	expected := 4.3639
	result := r.Distance(lat1, lon1, lat2, lon2)

	assert.InDelta(t, expected, result, 0.001)
}
func TestDegToRad(t *testing.T) {
	r := services.Routing{}

	degree := 45.0
	expected := math.Pi / 4.0
	result := r.DegToRad(degree)

	if !floatEquals(result, expected) {
		t.Errorf("degToRad calculation incorrect. Got: %f, Expected: %f", result, expected)
	}

}

func floatEquals(a, b float64) bool {
	const tolerance = 1e-6
	diff := math.Abs(a - b)
	max := math.Max(math.Abs(a), math.Abs(b))
	return diff <= tolerance*max
}
