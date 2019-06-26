package main

import (
	"fmt"
	"flag"
)

//flags
var (
	args = flag.String("a", "Pb qdph lv pqwnb.", "-a option is give cipher")
)

func main() {
	flag.Parse()
	//shift 1~25 word
	for i := 1; i < 26; i++ {
		fmt.Printf("----Shift %v word----\n", i)
		for _, s := range *args {
			if (s >= 65 && s <= 90) {
				s += int32(i)
				if s > 90 {
					s = 65 + (s - 91)
				}
			}else {
				s += int32(i)
				if (s - int32(i)) < 65 {
					s -= int32(i)
				}else if s > 122 {
					s = 97 + (s - 123)
				}
			}
			fmt.Print(string(s))
		}
		fmt.Printf("\n\n")
	}
}
