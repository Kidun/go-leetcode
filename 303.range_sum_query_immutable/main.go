package SumRange

type NumArray struct {
	a []int
}

func Constructor(nums []int) NumArray {
	// we store precalculated sums of all previous elements
	res := NumArray{}
	sum := 0
	for _, num := range nums {
		res.a = append(res.a, sum+num)
		sum += num
	}
	return res
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > 0 {
		return this.a[right] - this.a[left-1]
	} else {
		return this.a[right]
	}
}
