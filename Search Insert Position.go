func searchInsert(nums []int, target int) int {
    l := 0
    r := len(nums) - 1
    idx := 0
    if target > nums[r]{
        return len(nums)
    }else if target < nums[l]{
        return 0
    }else{
        for l < r{
            idx = (l+r)/2
            if nums[idx] == target{
                return idx
            }else if target > nums[idx]{
                l++
            }else{
                r--
            }
        }
    }    
    return l
}
