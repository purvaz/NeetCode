package BinarySearch

func SearchMatrix(matrix [][]int, target int) bool {

	startArr, endArr := 0, len(matrix)-1
	midArr := -1
	for startArr <= endArr {
		midArr = (startArr + endArr) / 2
		if target >= matrix[midArr][0] && target <= matrix[midArr][len(matrix[midArr])-1] {
			break
		} else if target >= matrix[startArr][0] && target < matrix[midArr][0] {
			endArr = midArr - 1
		} else {
			startArr = midArr + 1
		}
	}

	start, end := 0, len(matrix[midArr])-1

	if target == matrix[midArr][start] || target == matrix[midArr][end] {
		return true
	}
	for start <= end {
		mid := (start + end) / 2
		if matrix[midArr][mid] == target {
			return true
		} else if target >= matrix[midArr][start] && target < matrix[midArr][mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}
