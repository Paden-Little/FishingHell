package fisher

import (
	"log"
	"main/datalayer"
	"main/protobufs"
	"math"
	"math/rand"
	"time"
)

func GetFish(pool string, poleID int32, baitID int32) *protobufs.Fish {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := pickFish(pool, baitID, rng)
	fishType := datalayer.GetFish(id)
	weight := generateWeight(fishType.Low_weight, fishType.High_weight, fishType.Peak, fishType.Spread, rng)
	fish := &protobufs.Fish{
		Id:         int32(fishType.ID),
		Name:       fishType.Name,
		Taxonomy:   fishType.Taxonomy,
		Pool:       pool,
		Weight:     float32(weight),
		TimeCaught: time.Now().Format("2006-01-02 15:04:05"),
		BaitUsed:   baitID,
	}

	return fish
}

func probabilityDensity(weight float64, peak float64, spread float64) float64 {
	return math.Exp(-(math.Pow((weight-peak), 2) / (2 * math.Pow(spread, 2))))
}

func uniform(rng *rand.Rand, a, b float64) float64 {
	return a + rng.Float64()*(b-a)
}

func generateWeight(low_weight float64, high_weight float64, peak float64, spread float64, rng *rand.Rand) float64 {
	var weight float64
	spread = (high_weight - low_weight) / spread

	for {
		gen := uniform(rng, low_weight, high_weight)
		probability := probabilityDensity(gen, peak, spread)

		if rng.Float64() < probability {
			weight = gen
			break
		}
	}

	return weight
}

func pickFish(poolID string, baitID int32, rng *rand.Rand) int {
	catchValues, err := datalayer.GetCatchValues(poolID, baitID)
	if err != nil {
		log.Fatalf("Error fetching catch values: %v", err)
	}

	catches := make([]int, 0)
	for n := range catchValues {
		for range catchValues[n].Value {
			catches = append(catches, catchValues[n].FishID)
		}
	}

	if len(catches) == 0 {
		log.Fatalf("No fish available for poolID %s and baitID %d", poolID, baitID)
	}

	return catches[rng.Intn(len(catches))]
}
