package account

import "fmt"

func ListRegion() {
	for k, v := range getRegionList() {
		fmt.Println(k, v)
	}
}
