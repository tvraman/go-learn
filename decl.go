package main

import "fmt"

var i, j int = 1, 2
var c, python, java = true, false, "no!"
func local() {
    var i, j int = 1, 2
    k := 3
    c, python, java := true, false, "no!"

    fmt.Println("Local: ",  i, j, k, c, python, java)
}

func main() {
    fmt.Println(i, j, c, python, java)
	local()
}
