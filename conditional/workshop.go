package conditional

import "fmt"

func FindMax() {
	a := 34
	b := 35
	c := 33

	if a > b && a > c {
		fmt.Println("a en buyuk")
	} else if b > a && b > c {
		fmt.Println("b en buyuk")
	} else if c > a && c > b {
		fmt.Println("c en buyuk")
	}
}
