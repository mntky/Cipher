package main

import (
	"fmt"
)


func main() {
	//My name is mntky. hi
	// ->3
	//Pb qdph lv pqwnb. kl
	test := "Pb qdph lv pqwnb. kl"

	for _, s := range test {
		fmt.Printf("%v,", s)
		s += 3
		fmt.Printf("%v\n", s)
	}
}
