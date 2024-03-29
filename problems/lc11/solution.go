package lc11

// 11. ContainerWith Most Water

// You are given an integer array height of length n.
// There are n vertical lines drawn such that the
// two endpoints of the ith line are (i, 0) and (i, height[i]).
//
// Find two lines that together with the x-axis form a container,
// such that the container contains the most water.
//
// Return the maximum amount of water a container can store.
//
// Notice that you may not slant the container.

// Constraints:
//
// n == height.length
// 2 <= n <= 105
// 0 <= height[i] <= 104

func maxArea(height []int) int {
	var (
		rightIdx        = len(height) - 1
		leftIdx         int
		area            int
		availableHeight int
		availableWidth  int
	)

	for leftIdx < rightIdx {
		availableHeight = min(height[leftIdx], height[rightIdx])
		availableWidth = rightIdx - leftIdx
		area = max(area, availableHeight*availableWidth)

		if height[leftIdx] < height[rightIdx] {
			leftIdx++
		} else {
			rightIdx--
		}
	}
	return area
}
