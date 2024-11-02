package practice

import (
	"fmt"
	"time"
)

func StartSwitchPractice() {
	normalSwitch()
	initValueSwitch()
	comboSwitch()
}

func normalSwitch() {
	fmt.Sprintln("--------------一般switch練習------------")
	date := time.Monday
	switch date {
	case time.Monday:
		fmt.Println("星期一")
	case time.Tuesday:
		fmt.Println("星期二")
	case time.Wednesday:
		fmt.Println("星期三")
	case time.Thursday:
		fmt.Println("星期四")
	case time.Friday:
		fmt.Println("星期五")
	}
}

func initValueSwitch() {
	fmt.Println("--------------賦予值與複雜判斷式------------")
	switch date := time.Saturday; {
	case date == time.Monday || date == time.Tuesday || date == time.Wednesday || date == time.Thursday || date == time.Friday:
		fmt.Println("是平日")
	default:
		fmt.Println("是假日")
	}
}

func comboSwitch() {
	fmt.Println("-------------針對值得多個判斷-------------")
	switch age := 32; age {
	case 22, 23, 24:
		fmt.Println("都是2開頭")
	case 33, 32, 35:
		fmt.Println("都是3開頭")
	}
}
