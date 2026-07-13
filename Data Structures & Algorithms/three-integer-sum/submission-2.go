func threeSum(nums []int) [][]int {
    sort.Ints(nums)

    res := [][]int{}

    n := len(nums)

    for i := 0; i < n-2; i++ {
        
        // skip duplicate first element
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        left := i + 1
        right := n-1

        for left < right {
            sum := nums[i] + nums[left] + nums[right]

            if sum < 0 {
                left++
            } else if sum > 0 {
                right--
            }else {
                res = append(res, []int{nums[i], nums[left], nums[right]})
                left++
                right--

                // skip duplicate on left
                for left < right && nums[left] == nums[left-1] {
                    left++
                }

                // skip duplicate on right
                for left < right && nums[right] == nums[right+1] {
                    right--
                }
            }
        }
    }

    return res
}
