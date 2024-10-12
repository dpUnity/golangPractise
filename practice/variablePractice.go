package practice

import "fmt"

// 宣告方式 var 變數名稱 型別 ＝ 值
var name string = "Koli"

// 一次宣告多個變數
var (
	character string = "Wife"
	age       string = "32"
	marryDate string = "2021/02/16"
)

func Start() {
	var person string = "cute"
	fmt.Println("單一宣告結果:", name, person)
	fmt.Println("一次宣告多個:", character, age, marryDate)
}
