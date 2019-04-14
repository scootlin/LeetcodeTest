var list = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t","u","v"},
    "9": []string{"w", "x","y", "z"},
}
func letterCombinations(digits string) []string {
    var rootSlice []string   
	if len(digits) <= 1 {
		if _, ok := list[digits]; ok {
			return list[digits]
		} else {
			return []string{}
		}
	}
	ch := strings.Split(digits, "")
	rootSlice = append(rootSlice, list[ch[0]]...)
	for y := 1; y < len(ch); y++ {
		tmp := []string{}
		roolen := len(rootSlice)
		chs1 := list[ch[y]]
		for x := 0; x < roolen; x++ {
			// fmt.Println(rootSlice)
			for x1 := 0; x1 < len(chs1); x1++ {
				str := rootSlice[x] + chs1[x1]
				tmp = append(tmp, str)
			}
		}
		rootSlice = tmp
	}
	return rootSlice
}
