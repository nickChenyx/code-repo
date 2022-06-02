package util

import "fmt" 

func IsTest(t *int) bool {
    fmt.Println("util.isTestCall")
    if t == nil {
        panic("t can't be nil")
    }
    return *t == 0
}
