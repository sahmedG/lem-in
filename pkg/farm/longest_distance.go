package farm

func (farm *Farm) Longest_Distance() int {
	max := -9999

	for _, distance := range farm.distances {
		if distance > max {
			max = distance
		}
	}

	return max
}

func (farm *Farm) Shortest_Distance() int {
	min := 9999

	for _, distance := range farm.distances {
		if distance < min {
			min = distance
		}
	}

	return min
}
