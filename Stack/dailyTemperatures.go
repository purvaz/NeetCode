package Stack

type TempStruct struct {
	index int
	value int
}

func DailyTemperatures(temperatures []int) []int {

	stack := []TempStruct{}
	var result = make([]int, len(temperatures))

	for i, v := range temperatures {
		for len(stack) > 0 && v > stack[len(stack)-1].value {
			index := stack[len(stack)-1].index
			stack = stack[:len(stack)-1]
			result[index] = i - index
		}
		stack = append(stack, TempStruct{index: i, value: v})
	}
	return result
}
