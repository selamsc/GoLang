package loops

import "fmt"

func Asalmi() {
	input := 0
	fmt.Println("bir sayi giriniz")
	fmt.Scanln(&input)

	isAsal := true
	for i := 2; i < input; i++ {
		if input%i == 0 {
			isAsal = false
		}
	}
	if !isAsal {
		fmt.Println("girdiginiz sayi asal degil")
	} else {
		fmt.Println("girdiginiz sayi asal")
	}

}
