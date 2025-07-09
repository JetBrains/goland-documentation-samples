package main

import (
	"fmt"
	"nil_dereference_analysis/common"
)

func main() {
	user := common.NewUser("", 21)
	fmt.Printf("Age: %d, %s", user.Age, user.Name)
}
