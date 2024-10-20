package practice

import (
	"fmt"
	"time"
)

// 宣告方式 var 變數名稱 型別 ＝ 值
var name string = "Koli"

// 一次宣告多個變數
var (
	character  string    = "Wife"
	isFemale   bool      = true
	updateTime time.Time = time.Now()
)

// 可省略型別和初始值
var (
	omitCharacter  = "Wife"
	omitIsFemale   bool //按書上來說預設值會是zero value,但這邊是false
	omitUpdateTime = time.Now()
)

func Start() {
	declarPractice()
	omitPractice()
}

func declarPractice() {
	var person string = "cute"
	fmt.Println("單一宣告結果:", name, person)
	fmt.Println("一次宣告多個:", character, isFemale, updateTime)
}

func omitPractice() {
	fmt.Println("省略練習結果:", omitCharacter, omitIsFemale, omitUpdateTime)
}
