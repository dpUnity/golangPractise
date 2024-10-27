package practice

import (
	"fmt"
	"time"
)

func StartPointersPractice() {
	variablePointer()
	obtainValueFormPointer()
}

func variablePointer() {
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

func obtainValueFormPointer() {
	fmt.Println("--------------指標取值練習--------------")
	var count1 *int
	count2 := new(int)
	countTemp := 6
	count3 := &countTemp
	nowTime := time.Now()
	tempTime := &nowTime

	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	}

	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}

	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}

	if tempTime != nil {
		fmt.Printf("time: %#v\n", *tempTime)
		fmt.Printf("time round two: %#v\n", tempTime.String())
	}
}
