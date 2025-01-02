package main

import "fmt"

// add 函数接收两个整数参数，并返回它们的和
func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(5, 3)
    fmt.Printf("5 + 3 = %d\n", sum)
}
