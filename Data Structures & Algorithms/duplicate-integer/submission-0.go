func hasDuplicate(nums []int) bool {
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        _, ok := m[nums[i]]
        if ok {
            return true
        }
        m[nums[i]]++
    }

    return false
}
