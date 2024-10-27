package practice

import (
	"errors"
	"fmt"
)

func StartIfPractice() {
	fmt.Println("--------IF 起始賦值練習--------")
	num := 22
	if error := validate(num); error == nil {
		fmt.Printf("輸入的數字是：%v", num)
	} else {
		fmt.Println("輸入的值有問題 ： " + error.Error())
	}
}

func validate(num int) error {
	var errorCode error = nil
	if num > 100 {
		errorCode = errors.New("不能超過100")
	} else if num%7 == 0 {
		errorCode = errors.New("不能是7的倍數")
	} else if num < 0 {
		errorCode = errors.New("不能是負數")
	}

	return errorCode
}
