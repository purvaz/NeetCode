package Stack

import (
	"strconv"
)

func EvalRPN(tokens []string) int {

	opStack := []int{}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			op2 := pop(&opStack)
			op1 := pop(&opStack)
			opStack = append(opStack, getResult(op1, op2, tokens[i]))
		} else {
			num, _ := strconv.Atoi(tokens[i])
			opStack = append(opStack, num)
		}
	}
	return opStack[0]
}

func pop(stack *[]int) int {
	n := len(*stack)
	num := (*stack)[n-1]
	*stack = (*stack)[:n-1]
	return num
}

func getResult(op1, op2 int, operator string) int {

	switch operator {
	case "+":
		return op1 + op2
	case "-":
		return op1 - op2
	case "*":
		return op1 * op2
	case "/":
		return op1 / op2
	default:
		return 0
	}
}
