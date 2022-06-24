package conditional

import "fmt"

func Cond1() {
	balance := 400
	withdraw := 300
	if withdraw > balance {
		fmt.Println("Account has no enough balance")
	} else if balance == withdraw {
		fmt.Println("paran bitti")
	}
	balance = balance - withdraw
	fmt.Println("successfully withdrawed, last balance : " + fmt.Sprint(balance))
}
