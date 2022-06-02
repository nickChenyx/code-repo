package main

import (
	_ "example.com/import_panic_init/a"
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
