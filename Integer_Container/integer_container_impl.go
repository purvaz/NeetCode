package integercontainer

import "sort"

type IntegerContainerImpl struct {
	AbstractIntegerContainer
	values []int
}

func NewIntegerContainerImpl() *IntegerContainerImpl {
	return &IntegerContainerImpl{
		AbstractIntegerContainer: AbstractIntegerContainer{},
		values:                   make([]int, 0),
	}
}

func (i *IntegerContainerImpl) Add(value int) int {
	i.values = append(i.values, value)
	return len(i.values)
}

func (i *IntegerContainerImpl) Delete(value int) bool {

	for j := 0; j < len(i.values); j++ {
		if i.values[j] == value {
			i.values = append(i.values[:j], i.values[j+1:]...)
			return true
		}
	}
	return false
}

func (i *IntegerContainerImpl) GetMedian() *int {

	if len(i.values) == 0 {
		return nil
	}

	sorted := make([]int, 0, len(i.values))
	for _, value := range i.values {
		sorted = append(sorted, value)
	}

	sort.Ints(sorted)
	mid := len(sorted) / 2

	if len(sorted)%2 == 0 {
		return &sorted[mid-1]
	} else {
		return &sorted[mid]
	}
}
