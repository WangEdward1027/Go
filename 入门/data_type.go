package main

import "fmt"

func printBool(b bool){
    fmt.Println(b)
}

func printInts(a int, b uint8){
    fmt.Println(a, b)
}

func printFloat(f float64){
    fmt.Println(f)
}

func printComplex(c complex128){
    fmt.Println(c);
    fmt.Println(real(c))
    fmt.Println(imag(c))
}

func printString(s string){
    fmt.Println(s)
    fmt.Println(len(s))
}


func main(){
    var isReady bool = true

    var a int = 100
    var b uint8 = 255

    var pi float64 = 3.14159265358979

    var c complex128 = 1 + 2i

    var s string = "Hello, 世界"

    printBool(isReady)
    printInts(a, b)
    printFloat(pi)
    printComplex(c)
    printString(s)
}
