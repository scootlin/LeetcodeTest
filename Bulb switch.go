func main() {
    list := randlist(10)
    fmt.Println("List => ",list)
    fmt.Println(moment(list))
}

func randlist(max int)[]int{
    list:=[]int{}
    match := make(map[int]int,max)
    size := max
    rand.Seed(time.Now().UnixNano())
	for size > 0 {
		x := rand.Intn(max+1)
        if _,ok := match[x]; !ok &&  x != 0{
            list = append(list,x)
            size --
            match[x]=size
        }
	}
    return list
}

func moment(list []int)int{
    count := 0
    switcher := make(map[int]bool,len(list)+1)
    switcher[0]=true
    for _,val:=range list{      
        on := true
        for i:=1; i<val;i++{
            if !switcher[i]{
                on=false
                break
            }
        }
        if on {
            count++
        }
        switcher[val] = true
    }
    return count
}

