package Numbers

import (
	"fmt"
	"math/rand"
)

func getWeightedStrings(strings []string, weights []float64) string {

	randomNumber := rand.Float64()

	cumulativeProb := make([]float64, len(weights))

	cumulativeProb[0] = weights[0]
	for i := 1; i < len(weights); i++ {
		cumulativeProb[i] = cumulativeProb[i-1] + weights[i]
	}

	for index, value := range cumulativeProb {
		if randomNumber < value {
			return strings[index]
		}
	}

	return strings[len(strings)-1]

}

func CallWeigtedStrings(strings []string, weights []float64, count int) {

	occurrenceMap := make(map[string]int, len(strings))

	for i := 0; i < count; i++ {
		occurrenceMap[getWeightedStrings(strings, weights)]++
	}

	for str, count := range occurrenceMap {
		fmt.Println("String ", str, " occurs ", count, " times. ")
	}
}
