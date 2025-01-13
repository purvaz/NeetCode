package Stack

import (
	"sort"
)

type Car struct {
	position int
	time     int
}

func CarFleet(target int, position []int, speed []int) int {

	stack := []float32{}
	// Combine into 1 array
	pair := make([]Car, len(position))
	for index, value := range position {
		pair[index] = Car{value, speed[index]}
	}

	// Sort the array based on position
	sort.Slice(pair, func(i, j int) bool {
		return pair[i].position < pair[j].position
	})
	// fmt.Println(pair)

	for i := len(pair) - 1; i >= 0; i-- {
		stack = append(stack, (float32(target-pair[i].position) / float32(pair[i].time)))
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack)
}
