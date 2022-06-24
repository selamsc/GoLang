package loops

import "fmt"

func LoopFunc() {
	text := "loop dongusu"
	i := 1

	for i <= 5 {
		fmt.Println(text)
		i = i + 1
	}
}
