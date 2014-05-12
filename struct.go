package main
import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println("Point: ", v)
    v.X = 4
	fmt.Println("New Point: ", v)
}
