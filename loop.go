package main

import "fmt"

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println("Sum: ", sum)
	
i := 1
    for  i < 1000 {
        i += i
    }
	fmt.Println("i: ", i)
}
