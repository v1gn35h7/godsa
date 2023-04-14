package algo

type KnapsackItem struct {
	P, W int
}

func FractionalKanpsack(w int, list []KnapsackItem) int {
	profit := 0
	weight := 0
	for _, v := range list {
		if weight >= w {
			break
		}

		if (v.W + weight) < w {
			profit = profit + v.P
			weight = weight + v.W
		} else {
			rw := w - weight
			frac := float64(rw) / float64(v.W)
			profit = profit + int(float64(v.P)*frac)
			weight = weight + int(float64(v.W)*frac)
		}
	}

	return profit
}
