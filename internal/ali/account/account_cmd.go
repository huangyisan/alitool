package account

import "fmt"

func ListAccount() {
	fmt.Printf("Account List:\n\n")
	for k, _ := range getAccountMap() {
		fmt.Printf("%s\n", k)
	}
}
