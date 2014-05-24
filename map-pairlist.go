package main

import (
	"fmt"
	"github.com/tvraman/go-learn/utils"
)

// Exercise sortmap.SortedKeys
func main() {
	m := map[string]int{
		"a": 100,
		"b": 7,
		"c": 11,
	}

	for key, value := range m {
		fmt.Printf("%v:   %v\n", key, value)
	}
	fmt.Println("Sorted")
	p := pairlist.SortMapByValue(m)
	fmt.Println(p)
	for i := range p {
		fmt.Println(p[i].key)
	}
}
