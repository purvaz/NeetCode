package Stack

import "strings"

func GenerateParenthesis(n int) []string {

	var stack = []string{}
	var res = []string{}

	var backtrack func(int, int)
	backtrack = func(openCount, closeCount int) {

		if openCount == closeCount && openCount == n && closeCount == n {
			res = append(res, strings.Join(stack, ""))
			return
		}

		if openCount < n {
			stack = append(stack, "(")
			backtrack(openCount+1, closeCount)
			popStack(&stack)
		}

		if closeCount < openCount {
			stack = append(stack, ")")
			backtrack(openCount, closeCount+1)
			popStack(&stack)
		}
	}
	backtrack(0, 0)
	return res
}

func popStack(stack *[]string) {
	length := len(*stack)
	*stack = (*stack)[:length-1]
}
