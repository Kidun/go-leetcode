package main

func largestRectangleArea(heights []int) int {
	maxArea := 0

	var max = func(x1 int, x2 int) int {
		if x1 > x2 {
			return x1
		}
		return x2
	}

	// Prepare a stack for [index, height]
	stack := [][2]int{}

	// Iterate through all indices and heights
	for i, h := range heights {
		start := i

		// While stack is not empty and last item on stack
		// is greater than current height
		for len(stack) > 0 && stack[len(stack)-1][1] > h {
			// pop the last item from stack
			index := stack[len(stack)-1][0]
			height := stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]

			maxArea = max(maxArea, height*(i-index))

			// Set start to last popped index
			start = index
		}

		// Append start and height to stack
		stack = append(stack, [2]int{start, h})
	}

	// calculate area for all items remaining in stack
	for _, item := range stack {
		i := item[0]
		h := item[1]
		maxArea = max(maxArea, h*(len(heights)-i))
	}
	return maxArea
}

/*
	type Rectangle struct {
	    Index int
	    Height int
	}

	func largestRectangleArea(heights []int) int {
	    res := 0
	    stack := make([]Rectangle, len(heights))

	    for i, h := range heights {
	        startIndex := i
	        for len(stack) > 0 && stack[len(stack)-1].Height >= h {
	            startIndex = stack[len(stack)-1].Index
	            res = max(res, stack[len(stack)-1].Height * (i - stack[len(stack)-1].Index))
	            stack = stack[:len(stack)-1]
	        }
	        stack = append(stack, Rectangle{startIndex, h})
	    }

	    for len(stack) > 0 {
	        res = max(res, stack[len(stack)-1].Height * (len(heights) - stack[len(stack)-1].Index))
	        stack = stack[:len(stack)-1]
	    }
	    return res
	}
*/

func main() {

}
