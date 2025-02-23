package practice

import "fmt"

func StartArrayPractice() {
	fmt.Println("-----------陣列練習開始-----------")
	declareArray()
	compareArray()
}

func declareArray() {
	var ary1 [10]int
	fmt.Println(ary1)

	var ary2 = [10]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ary2)

	var ary3 = [10]int{1, 3: 4, 6, 7: 9}
	fmt.Println(ary3)
}

func compareArray() {
	var ary1 = [...]int{0, 1, 2, 3, 4, 5}
	var ary2 [6]int
	var ary3 = [7]int{0, 1, 2, 3, 4, 5, 6}
	var ary4 = [...]int{0, 0, 0, 0, 0, 0}

	fmt.Printf("ary1 == ary2 : %v\n", ary1 == ary2)
	fmt.Printf("ary3 : %v\n", ary3)
	// fmt.Printf("ary1 == ary3 : ", ary1 == ary3) // 因長度不同所以會直接出錯
	fmt.Printf("ary2 == ary4 : %v\n", ary2 == ary4)
}
