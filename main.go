package main

import (
    "fmt"
    golib "github.com/stamhe/golib"
)

func main() {
    var test1 string
    var test2 string
    
    test1 = golib.ReturnStr()
    test2 = golib.ReturnStr2()
    
    fmt.Printf("ReturnStr from pack1: %s\n", test1)
    fmt.Printf("Integer from pack1: %d\n", golib.Pack1Int)
    fmt.Printf("Float from pack1: %.2f\n", golib.Pack1Float)
    
    fmt.Println()
    
    fmt.Printf("ReturnStr from pack2: %s\n", test2)
    fmt.Printf("Integer from pack2: %d\n", golib.Pack2Int)
    fmt.Printf("Float from pack2: %.2f\n", golib.Pack2Float)
}
