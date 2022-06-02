package panic_init

import "fmt"
import "util"

func init() {
    util.IsTest(nil)
    fmt.Println("panic_init run..")
}
