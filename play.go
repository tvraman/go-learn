package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    fmt.Println("Welcome to the playground!")
	fmt.Println("Running this locally, so no sandbox!")
    fmt.Println("The time is", time.Now())

    fmt.Println("And if you try to open a file:")
    fmt.Println(os.Open("play.go"))

    fmt.Println("Or access the network:")
    fmt.Println(net.Dial("tcp", "google.com:80"))
}
