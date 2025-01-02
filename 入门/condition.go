package main

import "fmt"

func main(){
    var age int = 25
    fmt.Println("age = ", age)

    if age < 18{
        fmt.Println("未成年")
    }else if age <= 30{
        fmt.Println("青年")
    }else if age <= 60{
        fmt.Println("中年")
    }else{
        fmt.Println("老年")
    }
}
