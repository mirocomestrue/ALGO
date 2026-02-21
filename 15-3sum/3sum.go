import (
    "fmt"
    "sort"
)
func threeSum(nums []int) [][]int {
    var ret [][]int
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
            if n1+n2+n3 >= 0 {
                if n1+n2+n3 == 0 {
                    ret = append(ret, []int{n1, n2, n3})
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