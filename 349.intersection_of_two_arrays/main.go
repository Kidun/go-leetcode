package intersection

func intersection(nums1 []int, nums2 []int) []int {
	// better having in hashmap the largest slice
	if len(nums2) > len(nums1) {
		return intersection(nums2, nums1)
	}
	// store slice in hashmap for O(1) search in it
	nums1Map := map[int]bool{}
	for _, num := range nums1 {
		nums1Map[num] = true
	}
	res := []int{}
	// parse 2nd slice, checking for presence in the 1st slice
	for _, num := range nums2 {
		if nums1Map[num] {
			res = append(res, num)
			nums1Map[num] = false
		}
	}
	return res
}
