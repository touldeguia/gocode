package main 

import (
    "fmt"
    "os"
)

func main() { 
    s, sep := "", ""
    for _, arg := range os.Args[1:] { 
        s += sep + arg 

    } 
    fmt.Println(s)
    fmt.Println(os.Args[1:])
} 

