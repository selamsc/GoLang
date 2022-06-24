package loops

import "fmt"

func Tahmin1() {
	belirtec := 46
	tahmin := 0

	fmt.Println("tahmin et")
	fmt.Scanln(&tahmin)
	for tahmin != belirtec {
		if tahmin > belirtec {
			fmt.Println("Buyuk deger girdiniz")
			fmt.Scanln(&tahmin)
		}
		if tahmin < belirtec {
			fmt.Println("Kucuk deger girdiniz")
			fmt.Scanln(&tahmin)
		}
	}

	fmt.Printf("dogru tahmin ettiniz: %v", belirtec)

}
