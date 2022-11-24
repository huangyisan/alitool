package account

import "fmt"

func DoListAccount() {
	fmt.Printf("Account List:\n\n")
	for k, _ := range accountMap {
		fmt.Printf("%s\n", k)
	}
}
