package fisher

import (
	"math/rand"
	"testing"
	"time"
)

func TestGenerateFishWeightCustom(t *testing.T) {
	//Default peak is (low + high) / 2
	//Default spread is (high - low) / 6, spread = stddev
	//The 6 is standard deviation but can be changed
	lowWeight := 1.0
	highWeight := 10.0
	peak := 5.0
	spread := 1.5

	weight := generateWeight(lowWeight, highWeight, peak, spread, rand.New(rand.NewSource(time.Now().UnixNano())))

	t.Logf("Generated Fish Weight: %.2f", weight)

	if weight < lowWeight || weight > highWeight {
		t.Errorf("Generated weight %.2f is out of range [%.2f, %.2f]", weight, lowWeight, highWeight)
	}
}
