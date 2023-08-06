package main


import "fmt"


func slices(){
	var s []string
	fmt.Println(s, len(s), s == nil)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0]="a"
	s[1]="b"
	s[2]="c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
    fmt.Println("len after set:", len(s))
    
    s = append(s, "d")
    s = append(s, "e", "f")

    fmt.Println("apd", s)

    c := make([]string, len(s))
    copy(c,s)

    fmt.Println("copy:", c)

    l := c[2:5]
    fmt.Println("l:", l, "and length:", len(l))

    l_two := c[2:]

    fmt.Println("l_two", l_two, len(l_two))

    l_five := c[:5]
    fmt.Println("l_five",l_five, len(l_five))

    // multi dimentional array
    twoD := make([][]int, 3)

    for i:=0; i<3; i++ {
    	inner_len := i+1
    	twoD[i] = make([]int, inner_len)
    	for j:=0; j<inner_len; j++ {
    		twoD[i][j]=i+j
    	}
    	fmt.Println("2d:", twoD)
    }


}

// learn maps

func maps(){

	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2

	fmt.Println(m)

	fmt.Println(m["k1"])
	fmt.Println(m["k3"])
	fmt.Println(len(m))
	delete(m, "k2")
	fmt.Println(m)

	_, prs := m["k2"]
	fmt.Println(prs)

	n :=map[string]int{"k3":1, "k4": 7}
	fmt.Println(n)

}


func main() {
    // learn slices
    slices()

    // learn maps
    maps()
}


