package main

import (
	_ "a"
	"fmt"
	_ "panic_init"
	"util"

	"bou.ke/monkey"
)

func main() {
	monkey.Patch(util.IsTest, func(t *int) bool {
		return true
	})
	fmt.Println("main run...")
}
