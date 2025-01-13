package BinarySearch

type TimeMap struct {
	store map[string][]ValueStamp
}

type ValueStamp struct {
	Val  string
	Time int
}

func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]ValueStamp)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, exists := this.store[key]; !exists {
		this.store[key] = make([]ValueStamp, 0)
	}
	this.store[key] = append(this.store[key], ValueStamp{Val: value, Time: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {

	var values []ValueStamp
	if _, exists := this.store[key]; exists {
		values = this.store[key]
	} else {
		return ""
	}
	// fmt.Println(values)

	left, right := 0, len(values)-1
	res := ""
	for left <= right {
		mid := (left + right) / 2
		if values[mid].Time == timestamp {
			return values[mid].Val
		} else if values[mid].Time < timestamp {
			res = values[mid].Val
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
