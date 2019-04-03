func longestPalindrome(s string) string {
    list := make(map[string]int)
    strs := strings.Split(s,"")
    maxstr:=""
    maxlen := 0
    // slen := len(s)
    if len(s) < 1{
        return s
    }
    for  curr , val := range strs{
        if idx,  ok:= list[val]; ok{
            offset:= curr - idx + 1
            if offet % 2 == 0 {
                
                
            }
            if s[idx:offset/2] == s[offset/2:curr]
            
            if offset > maxlen{
                maxlen = offset
                maxstr = s[idx:curr+1]
            }
            list[val] = 
        }else{
            list[val] = curr
        }
    }
    // fmt.Println(list)
    if maxlen < 1{
        return string(s[0])
    }else{
        return maxstr
    }
}

func compares(str1, str2 string)bool{
    return str1 == str2
}