func twoSum(nums []int, target int) []int {
    numsmap := make( map[int]int)
    for i := range nums {
        numsmap[nums[i]] = i
    }
    
    for i := range nums {
        mapIndex, ok := numsmap[findNumber]
        
        if (ok && i != mapIndex) {
            return []int{i, mapIndex}
        }
    }
    return []int{}
}
