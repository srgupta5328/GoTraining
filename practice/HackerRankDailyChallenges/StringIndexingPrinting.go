package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, playground")
    s := "Helloplayground"
    s_1 := ""
    s_2 := ""
    for i := range s{
        if i % 2 ==0{
            s_1 += string(s[i])
        }else{
            s_2 += string(s[i])
        }
    }
    fmt.Printf("%s %s\n", s_1, s_2)
    
}