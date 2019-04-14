func strStr(haystack string, needle string) int {
    slen := len(needle)
    for i := 0 ; i < len(haystack)-slen+1 ; i++{
        str := haystack[i:i+slen]
        if string(str) == needle{
            return i
        }
    }
    return -1
}
