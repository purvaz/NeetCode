package Arrayshashing

import (
	"fmt"
	"math"
)

type KeyPosition struct {
	X int
	Y int
}

var keyboard = [][]string{
	{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"},
	{"a", "s", "d", "f", "g", "h", "j", "k", "l"},
	{"z", "x", "c", "v", "b", "n", "m", " "},
}

var keyPositionMap = make(map[string]KeyPosition, 0)

func CalculateKeyStrokes(word string) []int {

	returnArr := []int{}
	previousKey := KeyPosition{-1, -1}

	populateKeyPostitionMap(word)

	for _, char := range word {
		currentKey := keyPositionMap[string(char)]

		if previousKey.X == -1 && previousKey.Y == -1 {
			returnArr = append(returnArr, 0)
		} else {
			cost := math.Abs(float64(previousKey.X) - float64(currentKey.X) + float64(previousKey.Y) - float64(currentKey.Y))
			returnArr = append(returnArr, int(cost))
		}
		fmt.Println("Cost of ", string(char), " is ", returnArr[len(returnArr)-1])

		previousKey = currentKey
	}
	return returnArr
}

func populateKeyPostitionMap(word string) {

	for _, char := range word {
		if _, exists := keyPositionMap[string(char)]; !exists {
			currentKey := getKeyPosition(string(char))
			keyPositionMap[string(char)] = currentKey
		}
	}
}

func getKeyPosition(character string) KeyPosition {

	for row_index, r := range keyboard {
		for col_index, c := range r {
			if c == character || c == " " {
				return KeyPosition{X: row_index, Y: col_index}
			}
		}
	}
	return KeyPosition{X: -1, Y: -1}
}
