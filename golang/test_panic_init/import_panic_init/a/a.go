package a 

import "util"
import "bou.ke/monkey"
import "fmt"

func init() {
    monkey.Patch(util.IsTest, func(t *int) bool {
	return true
    })
    fmt.Printf("a init call util.IsTest: %v\n", util.IsTest(nil))
}
