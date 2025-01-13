package TwoPointers

// Approach 2:
/*
The 2 pointer approach executes in O(n) and doesnt require additional memory O(1).
Start with 2 vars maxL and maxR initialized to 0 and len-1 resp.
Compute min(maxL, maxR).
if maxL <= maxR, calculate min(maxL, maxR) - height[l]
	if the difference is > 0, add it to the total water.
	l++
else if maxL>maxR, calculate min(maxL, maxR) - height[r]
	if the difference is > 0, add it to the total water.
	r--
In the end, return water collected.
*/
func Trap(height []int) int {
	l, r := 0, len(height)-1
	maxL, maxR := height[0], height[len(height)-1]
	waterTrapped := 0
	for l <= r {
		unitsCollected := 0
		if maxL <= maxR {
			unitsCollected = min(maxL, maxR) - height[l]
			if unitsCollected > 0 {
				waterTrapped += unitsCollected
			}
			if height[l] > maxL {
				maxL = height[l]
			}
			l++
		} else {
			unitsCollected = min(maxL, maxR) - height[r]
			if unitsCollected > 0 {
				waterTrapped += unitsCollected
			}
			if height[r] > maxR {
				maxR = height[r]
			}
			r--
		}
	}
	return waterTrapped
}

// Approach 1:
/* Maintain 3 arrays: maxLeft, maxRight and minHeight that would store the respective values*/

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// func Trap(height []int) int {

// 	maxLeft := make([]int, len(height))
// 	maxRight := make([]int, len(height))
// 	waterCollected := 0

// 	maxLeft[0] = 0
// 	for i := 1; i < len(height); i++ {
// 		maxLeft[i] = max(maxLeft[i-1], height[i-1])
// 	}

// 	maxRight[len(height)-1] = 0
// 	for i := len(height) - 2; i >= 0; i-- {
// 		maxRight[i] = max(maxRight[i+1], height[i+1])
// 	}

// 	for i := 0; i < len(height); i++ {
// 		unit := ((min(maxLeft[i], maxRight[i])) - height[i])
// 		if unit > 0 {
// 			waterCollected += unit
// 		}
// 	}

// 	return waterCollected
// }
