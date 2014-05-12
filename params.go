package main

import "fmt"

func split(some int) (x, y int) {
	fmt.Println("Split: some is ", some)
    x = some * 4 / 9
	fmt.Println("x is now ", x)
    y = some - x
	fmt.Println("y is ", y)
    return
}

func main() {
    fmt.Println(split(17))
	fmt.Println(split(233))
}
