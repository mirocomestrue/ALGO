//go:build ignore
// +build ignore

package algo

func med(nums []int, idx float64) float64 {
	if float64(int(idx)) == idx {
		return float64(nums[int(idx)])
	}
	//fmt.Printf("%v %f\n", nums, float64(nums[int(idx)] + nums[int(idx) + 1]) / 2.0)
	return float64(nums[int(idx)]+nums[int(idx)+1]) / 2.0
}

func find(nums1 []int, nums2 []int, idx float64) float64 {
	//fmt.Printf("%v %v %f\n", nums1, nums2, idx)
	if len(nums1) == 0 {
		return med(nums2, idx)
	}
	if len(nums2) == 0 {
		return med(nums1, idx)
	}
	if idx == 0 {
		return float64(min(nums1[0], nums2[0]))
	}
	if idx == 0.5 {
		if nums1[0] < nums2[0] || (nums1[0] == nums2[0] && len(nums1) > 1 && len(nums2) > 1 && nums1[1] < nums2[1]) {
			if len(nums1) <= 1 {
				return float64(nums1[0]+nums2[0]) / 2.0
			}
			return float64(nums1[0]+min(nums1[1], nums2[0])) / 2.0
		}
		if len(nums2) <= 1 {
			return float64(nums2[0]+min(nums1[0])) / 2.0
		}
		return float64(nums2[0]+min(nums2[1], nums1[0])) / 2.0
	}
	// idx 번째 element 반환
	// idx 0 1 2 3
	// i.  0 0 1 1

	i := min(int((idx-1)/2.0), len(nums1)-1, len(nums2)-1) // i=1

	a, b := nums1[i], nums2[i]
	if a < b {
		// nums1[i] 까지는 다 셌다고 침
		return find(nums1[i+1:], nums2, idx-float64(i+1))
	} else if b < a {
		return find(nums1, nums2[i+1:], idx-float64(i+1))
	}
	// a==b
	return find(nums1[1:], nums2, idx-1.0)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return find(nums1, nums2, float64(len(nums1)+len(nums2)-1)/2.0)
}
