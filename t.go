package main

import "fmt"

func main() {
	type Worker struct {
		name string
	}
	type Key struct {
		out      string
		in       int
		optional string
		Worker
	}
	dict := map[Key]bool{
		{"A", 1, "yes", Worker{"w`"}}: true,
		{"A", 2, "no"}:                false,
		{"B", 2, "present"}:           false,
		{"B", 3, ""}:                  true,
	}
	elem, ok := dict[Key{"A", 1}]
	fmt.Println(elem)
	fmt.Println(ok)

	for key, value := range dict {
		fmt.Println(key, value)
	}

}
