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


func main() {
    // learn slices
    slices()
}


