package summaryRanges

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var make_range = func(start int, stop int) string {
		if start == stop {
			return strconv.Itoa(start)
		} else {
			return strconv.Itoa(start) + "->" + strconv.Itoa(stop)
		}
	}

	res := []string{}
	start_idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[start_idx] == i-start_idx {
			if i == len(nums)-1 {
				return append(res, make_range(nums[start_idx], nums[i]))
			}
			continue
		} else {
			res = append(res, make_range(nums[start_idx], nums[i-1]))
			if i == len(nums)-1 {
				return append(res, make_range(nums[i], nums[i]))
			}
			start_idx = i
		}
	}
	return res
}
