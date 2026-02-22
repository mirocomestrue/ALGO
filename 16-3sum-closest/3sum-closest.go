import (
    "fmt"
    "sort"
    "math"
)

func threeSumClosest(nums []int, target int) int {
    // 500*500*500 = 75,000,000
    // threeSum 과 같이 n * n(i, j는 투포인터)
    
    minAbs, ret := 100000, 0

    sort.Ints(nums)

    for k := len(nums)-1; k>1; k-- {
        if k<len(nums)-1 && nums[k] == nums[k+1] {
            continue
        }
        //fmt.Printf("%d\n", nums[k])
        i, j := 0, k-1
        for {
            if i >= j {
                break
            }
            n1, n2, n3 := nums[i], nums[j], nums[k]
            //fmt.Printf("%d %d,  %d %d,  %d %d\n", i, n1, j, n2, k, n3)
            abs := int(math.Abs(float64(n1+n2+n3-target)))
            if abs < minAbs {
                minAbs = abs
                ret = n1+n2+n3
            }
            if n1+n2+n3 >= target {
                if abs == 0 {
                    return target
                }
                // move j left
                for {
                    j = j-1
                    if j < 0 || nums[j] != nums[j+1] {
                        break
                    }
                }
                continue
            }
            // move i right
            for {
                i = i+1
                if i>=k || nums[i] != nums[i-1] {
                    break
                }
            }
        }
        
    }
    
    return ret
}