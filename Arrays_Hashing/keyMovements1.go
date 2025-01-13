package Arrayshashing

type Position struct {
	X int
	Y int
}

type KeyMovement struct {
	From Position
	To   Position
	Cost int
}

// var keyboard = [][]string{
// 	{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"},
// 	{"a", "s", "d", "f", "g", "h", "j", "k", "l"},
// 	{"z", "x", "c", "v", "b", "n", "m"},
// }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CalculateKeyMovements(word string) []KeyMovement {
	movements := make([]KeyMovement, 0, len(word))
	prevPos := Position{-1, -1}

	for _, char := range word {
		charPos := findCharPosition(string(char))
		if prevPos.X == -1 && prevPos.Y == -1 {
			// No previous position, starting movement
			movements = append(movements, KeyMovement{From: Position{-1, -1}, To: charPos, Cost: 0})
		} else {
			cost := abs(prevPos.X-charPos.X) + abs(prevPos.Y-charPos.Y)
			movements = append(movements, KeyMovement{From: prevPos, To: charPos, Cost: cost})
		}
		prevPos = charPos
	}
	return movements
}

func findCharPosition(char string) Position {
	for y, row := range keyboard {
		for x, key := range row {
			if key == char {
				return Position{X: x, Y: y}
			}
		}
	}
	return Position{-1, -1} // Not found
}
