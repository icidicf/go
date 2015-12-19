package main

import "fmt"

func main() {
	var m map[string]int
	m = make(map[string]int)
	m["route"] = 66
	i := m["route"]
	fmt.Printf("the i is %d\n", i)
}
