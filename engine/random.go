package engine

import "math/rand"

func RandFloat(randMin float64, randMax float64) float32 {
	return float32(rand.Float64()*(randMax-randMin) + randMin)
}
