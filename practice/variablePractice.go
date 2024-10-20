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
	declarePractice()
	omitPractice()
	shotDeclarePractice()
	oneLineDeclarePractice()
}

func declarePractice() {
	var person string = "cute"
	fmt.Println("--------------------變數練習--------------------")
	fmt.Println("單一宣告結果:", name, person)
	fmt.Println("一次宣告多個:", character, isFemale, updateTime)
}

func omitPractice() {
	fmt.Println("--------------------省略變數練習--------------------")
	//下面這樣會出錯，因為go會分不出seed是int還是int64，但rand.Seed需要的是int64
	// var seed = 123456789
	// rand.Seed(seed)
	fmt.Println("省略練習結果:", omitCharacter, omitIsFemale, omitUpdateTime)
}

func shotDeclarePractice() {
	fmt.Println("--------------------短宣告練習--------------------")
	shotCharacter := "wife"
	shotIsFemale := true

	fmt.Println("短宣告練習:", shotCharacter, shotIsFemale)

	shotName, shotAge, shotIsGay, shotUpdateTime := "Pao", 32, false, time.Now()
	fmt.Println("一次多個短宣告練習:", shotName, shotAge, shotIsGay, shotUpdateTime)

	shotFamilyCharacter, shotFamilyAge, shotUpdateFamilyTime := getFamily()
	fmt.Println("一次多個短宣告接回傳值練習:", shotFamilyCharacter, shotFamilyAge, shotUpdateFamilyTime)
}

func oneLineDeclarePractice() {
	fmt.Println("--------------------var單行練習--------------------")
	var name, character, nickName string = "Pao", "PaoPao", "PPPPao"
	fmt.Println("var單行宣告:", name, character, nickName)
	var name2, age, isFemale = "Koli", 32, true
	fmt.Println("var單行宣告不同型別:", name2, age, isFemale)
	var character2, age2, updateTime = getFamily()
	fmt.Println("var單行宣告接回傳值:", character2, age2, updateTime)

}

func getFamily() (string, int, time.Time) {
	return "Dad", 70, time.Now()
}
