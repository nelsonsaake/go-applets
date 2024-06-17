package data

import "math"

func BestValue(opts []Data) []Data {

	costValueRatio := []float64{}
	for _, opt := range opts {
		costValueRatio = append(costValueRatio, opt.MB/opt.GHC)
	}

	var maxCVR float64 = 0.0
	for _, cvr := range costValueRatio {
		maxCVR = math.Max(maxCVR, cvr)
	}

	bestValue := []Data{}
	for _, opt := range opts {
		if (opt.MB / opt.GHC) == maxCVR {
			bestValue = append(bestValue, opt)
		}
	}

	return bestValue
}
