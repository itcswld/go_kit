package main

import (
	"fmt"
)

func main() {
	s := " Tranquillity 記憶棉寧靜頸枕藍色TB212-BL"
	for _, v := range SliceChinglish(s, 17, 23) {
		fmt.Println(v)
	}
}
