package fisher

import (
	"math/rand"
	"testing"
	"time"
)

func TestGenerateFishWeightCustom(t *testing.T) {
	lowWeight := 1.0
	highWeight := 10.0
	peak := 0.0   // Leave peak and spread 0 for normal distribution
	spread := 0.0 // This will result in a default spread being calculated
	numSamples := 100

	// Create a new random generator with a seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Call the function to generate a fish weight
	weight := GenerateFishWeightCustom(rng, lowWeight, highWeight, peak, spread, numSamples)

	// Log the result using t.Logf
	t.Logf("Generated Fish Weight: %.2f", weight)

	// Optional: Check if the generated weight is within the expected range
	if weight < lowWeight || weight > highWeight {
		t.Errorf("Generated weight %.2f is out of range [%.2f, %.2f]", weight, lowWeight, highWeight)
	}
}

func TestWait(t *testing.T) {
	var ticks = Wait(0.0001)
	t.Log(ticks)
}
