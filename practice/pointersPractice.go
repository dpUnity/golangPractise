package practice

import (
	"fmt"
	"time"
)

func StartPointersPractice() {
	fmt.Println("--------------指標宣告練習--------------")
	var count1 *int
	fmt.Printf("count1: %#v\n", count1)
	count2 := new(int)
	fmt.Printf("count2: %#v\n", count2)
	countTemp := 5
	count3 := &countTemp
	fmt.Printf("count3: %#v\n", count3)
	time := &time.Timer{}
	fmt.Printf("time: %#v\n", time)
}
