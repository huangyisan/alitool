package account

import "fmt"

// ListAccount will print all AliAccount
func ListAccount() {
	fmt.Printf("Account List:\n\n")
	for k, _ := range getAccountMap() {
		fmt.Printf("%s\n", k)
	}
}
