package practice

import "fmt"

func StartRunePractice() {
	fmt.Println("-------------rune練習-----------")
	testWorld := "你好啊"
	testR := []rune(testWorld)
	for i := 0; i < len(testR); i++ {
		fmt.Println(string(testR[i]))
	}
	//也可以這樣做
	for _, value := range testWorld {
		fmt.Println(string(value))
	}
}
