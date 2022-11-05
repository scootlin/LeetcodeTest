func lengthOfLongestSubstring(s string) int {
    sLen := len(s)
    if sLen < 1 {
        return sLen
    }
    
    max := 0
    m1 := [256]int{}
    left := 0
    for right, val := range s {
        if m1[val] > 0 && left < m1[val] {
            left = m1[val]
        }
        m1[val] = right+1
        tempMax := right - left + 1
        if max < tempMax {
            max = tempMax
        }
    }
    return max
}
