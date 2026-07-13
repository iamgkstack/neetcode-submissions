func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    set := make(map[int]bool)

    for _, num := range nums {
        set[num] = true
    }


    longest := 0

    for num := range set {
        if !set[num-1] {
            current := num
            length := 1


            for set[current+1] {
                current++
                length++
            }

            if length > longest {
                longest = length
            }
        }
    }

    return longest
}
