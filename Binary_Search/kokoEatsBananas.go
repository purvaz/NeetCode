package BinarySearch

// import (
// 	"fmt"
// )

// func MinEatingSpeed(piles []int, h int) int {

// 	start, end := 1, Maximum(piles)
// 	result := end

// 	for start <= end {
// 		k := (start + end) / 2
// 		hours := 0
// 		for _, pile := range piles {
// 			// if pile/k <= 1 {
// 			// 	hours += 1
// 			// } else {
// 			// 	hours += pile / k
// 			// }
// 			hours += (pile + k - 1) / k
// 			fmt.Println(pile, k, (pile+k-1)/k)
// 		}
// 		fmt.Println("\nTotal hours: ", hours, "\n")

// 		if hours <= h {
// 			result = Min(result, hours)
// 			end = k - 1
// 		} else {
// 			start = k + 1
// 		}

// 	}
// 	return result
// }

// func Min(num1, num2 int) int {
// 	if num1 < num2 {
// 		return num1
// 	}
// 	return num2
// }

func Maximum(piles []int) int {
	max := -1
	for _, i := range piles {
		if i > max {
			max = i
		}
	}
	return max
}

// canEat returns true if all the bananas can be eaten within the time limit at speed
func canEat(piles []int, timeLimit, speed int) bool {
	timeNeed := 0
	for _, banana := range piles {
		timeNeed += (banana + speed - 1) / speed
		if timeNeed > timeLimit {
			return false
		}
	}

	return true
}

func MinEatingSpeed(piles []int, h int) int {
	// more bananas can be eaten with more speed
	// check if all the bananas can be eaten with a particular speed
	// if possible, store the ans & try smaller speed, greater speed otherwise
	lo, hi, ans := 1, Maximum(piles), 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if canEat(piles, h, mid) {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return ans
}
