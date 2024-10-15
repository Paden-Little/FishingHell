package fisher

import (
	"math"
	"math/rand"
	"time"
)

// customPDF calculates the custom probability density function (PDF)
func customPDF(x, peak, spread float64) float64 {
	return math.Exp(-((x - peak) * (x - peak)) / (2 * spread * spread))
}

// GenerateFishWeightCustom generates fish weight samples based on the custom PDF
func GenerateFishWeightCustom(rng *rand.Rand, lowWeight, highWeight, peak, spread float64, numSamples int) float64 {
	if peak == 0 {
		peak = (lowWeight + highWeight) / 2
	}
	if spread == 0 {
		spread = (highWeight - lowWeight) / 6
	}

	var weights []float64

	for i := 0; i < numSamples; i++ {
		weight := lowWeight + rng.Float64()*(highWeight-lowWeight)
		probability := customPDF(weight, peak, spread)
		if rng.Float64() < probability {
			weights = append(weights, weight)
		}
	}

	if len(weights) > 0 {
		return math.Round(weights[rng.Intn(len(weights))]*100) / 100
	}
	// Retry if no weights were accepted
	return GenerateFishWeightCustom(rng, lowWeight, highWeight, peak, spread, numSamples)
}

func Wait(successPercentage float64) int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	ticks := 0
	for {
		ticks++
		if rng.Float64() < successPercentage/100 {
			return ticks
		}
	}
}
