package Stack

import "fmt"

type Rectangle struct {
	index  int
	height int
}

func LargestRectangleArea(heights []int) int {

	maximum := 0
	stack := []Rectangle{}

	for index, height := range heights {
		prevIndex := index
		for len(stack) > 0 && stack[len(stack)-1].height > height {
			// calculate the area before popping
			calculatedHeight := stack[len(stack)-1].height * (index - stack[len(stack)-1].index)
			if maximum < calculatedHeight {
				maximum = calculatedHeight
			}
			prevIndex = stack[len(stack)-1].index
			// if the next height is lesser, pop previous
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, Rectangle{index: prevIndex, height: height})
	}

	for _, value := range stack {
		area := value.height * (len(heights) - value.index)
		fmt.Println("Computing area from stack ", area, value.height, len(heights), value.index)
		if area > maximum {
			maximum = area
		}
	}
	fmt.Println(stack)
	return maximum
}
