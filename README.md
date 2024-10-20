# golang練習

[TOC]

## 基本規則

* 同一個目錄下的go腳本一定要在同一個package下
* 同一個package可以不用import就直接用另一個腳本的方法
* 同專案不同package要引用的話路徑要這樣帶ex: {專案名稱}/{路徑}
* 一定要是main package才能使用main方法作為整個程式的進入點
* 不用分號



## 變數 

### 對應練習腳本

variablePractice

### 宣告方式

```go
var 變數名稱 型別 ＝ 值 //單一宣告方式
```

```go
var (
  變數名稱 型別 = 值
  變數名稱 型別 = 值
)//一次宣告多個
```



### 宣告注意事項

* 函式內宣告了變數一定要用

* 不管在哪宣告都可以省略型別，和TS一樣，語言會自行判斷型別

* 可以用英文外的語言命名，但拜託不要這麼做

* 按書上來說指定型別但不指定初始值，那預設是零值（zero value），可是bool卻是false，晚點看看零值到底是什麼

* 有時候一定要給型別不然類似型別會讓go混淆，ＥＸ：

  ```go
  var seed = 123456789
  rand.Seed(seed) //這樣會出錯，因為該方法要的是int64，但系統判斷seed為int
  ```



### 短宣告方式

```go
變數名稱 := 值

//可一次宣告多個
變數名稱1,變數名稱2,變數名稱N := 值1,值2,值N

//可用來接回傳多個變數的回傳值
變數名稱1,變數名稱2,變數名稱3 := getFamily()
func getFamily() (string, int, time.Time) {
	return "Dad", 70, time.Now()
}

```

### 短宣告注意事項

* 只能在方法內用喔！

### var單行多重宣告

* var 後面一次宣告多個變數名稱再接個型別，此方法會一次宣告多個同型別的變數

```go
var 變數名稱1,變數名稱2,變數名稱N 型別
```

* 若要不同型別，則不帶型別給相同數量的初始值

```go
var 變數名稱1,變數名稱2,變數名稱N = "Pao",32,true
```

* 拿來接多變數的回傳值也可以

```go
var 變數名稱1,變數名稱2,變數名稱3 = getFamily()
func getFamily() (string, int, time.Time) {
	return "Dad", 70, time.Now()
}
```

### 變數賦予值

* 直接等於就好

```go
var name = "Pao"
name = "Koli"
```

* 可以一次塞多個
```go
var name,age,isFemale = "Pao",32,false
name,age,isFemale = "Koli",31,true
```





## 運算符號

