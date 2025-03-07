# golang練習

[TOC]

## 基本規則

* 同一個目錄下的go腳本一定要在同一個package下
* 同一個package可以不用import就直接用另一個腳本的方法
* 同專案不同package要引用的話路徑要這樣帶ex: {專案名稱}/{路徑}
* 要呼叫別的package的方法用上面的範例import之後只要用路徑最後的腳本名稱.公開方法就好
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

* 按書上來說指定型別但不指定初始值，那預設是零值（zero value）

* 有時候一定要給型別不然類似型別會讓go混淆，ＥＸ：

  ```go
  var seed = 123456789
  rand.Seed(seed) //這樣會出錯，因為該方法要的是int64，但系統判斷seed為int
  ```



### 變數的零值

* 每個型別的零值不一樣，請參考下表

|類型|零值(預設值)|
|---|---|
|bool|false|
|所有數字|0|
|string|""(空字串)|
|其他物件類型的東西(存在heap的)|nil|



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
* 一定要賦予預設值

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

### 作用域

* 當全域變數和區域變數同名時，呼叫該名稱只會叫到區域變數的
* 完全沒辦法叫到全域變數的
* 要叫全域變數的只能另外開方法取得



## 運算符號

算術運算符（＋-*/%）和比較運算符（=、!、<、>）用法都跟Ｃ＃一樣



## 指標

### 對應練習腳本

pointersPractice

### 注意事項

* 當一個值被賦予指標，那他就只會在heap上，不會在stack上
* 指標的意思就是記憶體位置
* 同型別的指標互相比較只會是false，因為他比的是記憶體位置而不是裡面的值
* 結構體（struct）的指標印出來不會是明確的記憶體位置
* 結構體（struct）不用特地用＊取得指標指導的值，可以用String()印出值，透過＊只是在印出來時把＆拿掉而已，因為指標被解除了

### 宣告方式

1. 預設值為nil的宣告，此宣告方式因為沒帶該型別的預設值，所以預設為nil
  ```go
  var 變數名稱 *型別 //此方法不能用短宣告
  ```
2. 帶該型別預設值的宣告
  ```go
  var 變數名稱 new(型別)
  //方法內的話還可以這樣做↓
  變數名稱 := new(型別)
  ```
3. 取得該變數的指標的方法
  ```go
  變數 := &變數2
  ```

### 從指標取值

* 取得值的方式（依書上的說法）

```go
變數 := *指標變數
```

### 常數與列舉（enum）

* 常數的用法就只是把var改成const而已，其他規則都一樣
* go並沒有enum，直接用const來代替，只是可以用iota來協助宣告，ＥＸ：
	```go
	const (
		None = iota, // 0
		Empty				 // 1
		Normal			 // 2
		Medium			 // 3
		High				 // 4
		Full				 // 5
	)
	```



## 各種判斷式和迴圈

### IF

* 練習腳本：ifStatementPractice
* 宣告方式
	```go
	if 各種回傳boolean的東西 {
		功能
	}
	```

* 注意判斷式那邊不用加小括號
* 起始賦值的方式 （看起來就是for迴圈第一個分號前的東西）
	```go
	if 賦予值的程式;boolean判斷式 {
		功能
	}
	```

### switch

* 練習腳本：switchStatementPractice
* 不用break

* 也可以先賦予值再寫判斷
	```go
	switch 賦予值的程式;要判斷的值 {
		case 判斷式:
				程式碼區間
		case ...以此類推
	}
	```

* 可把所有判斷是塞在一個case裡面，這樣會吃同一個程式碼區間，就像C#連續多個case再接一個break
	```go
	a := 0
	switch a {
		case 0,1,2,3:
			fmt.Println("都是小於4的數字")
	} 
	```
	
* case那邊可以塞一個複雜的判斷式
	```go
	switch a := 4;{
		case a > 3 && a < 6:
			fmt.Println("a介於3到6之間")
	}
	```
	
	> * 要注意用複雜判斷式的時後，switch後面就不能擺要“判斷的值”，只能塞“宣告變數”或是不寫任何東西
	>
	> * 宣告後面一定要接;

### For迴圈

* 一般使用方式和C#一樣
	```go
	for 宣告值;判斷式;結束敘述{
		程式碼區間
	}
	```
	
* 什麼都不寫就是C#的wile(true)
	```go
	for {
		//就是無窮迴圈
	}
	```
	
* 只寫判斷是就是一般C#的wile
	```go
	a := 2
	for a < 1024 {
		a *= 2
	}
	```

* 搭配len可取得陣列或是切片的長度來跑for，就像C#的array.Length一樣

* 搭配range可以拿到任何集合的key、value上
	```go
	test := map[string]string 
	for key,value := range(test)
	```
	> value不接就不要寫
	> key不接就寫＿



## 	型別們
|型別|功能|備註|
|-|-|-|
|bool|就是true/false||
|int|有負數數字|大小長度取決於電腦是32還是64位元|
|int8|有負數數字|8位元|
|int16|有負數數字|16位元|
|int32|有負數數字|32位元|
|int64|有負數數字|64位元|
|uint|單純正數|大小長度取決於電腦是32還是64位|
|uint8|單純正數|8位元其實就是byte|
|uint16|單純正數|16位元|
|uint32|單純正數|32位元|
|uint64|單純正數|64位元|
|byte|單純正數|8位元，也等於uint8|
|float32|浮點數|32位元|
|float64|浮點數|64位元|
|string|就是字串|用`設值會是原始字串，不吃\n等。用"就會吃\n|

### 越界繞回

* go在發生溢味時，會出現越界繞回的動作，意思就是直接回到最小值，EX：
	```go
	var a byte = 255
	a += 1 //a會變成0
	```

* 要解決可以用math/big的套件

### String型別

* ``括住的字串會顯示原始內容，意思就是打換行就會換行，但用\n不會
* ""括住的字串會顯示轉譯過的內容，意思就是\n才會換行

### rune型別
* 練習腳本：runePractice
	
* uint32的別稱，同時可以拿來處理UTF-8的文字格式，不然用一般string來處理像是中文一樣的文字會印不出來，但是轉乘rune就可以，EX:
	
	```go
	name := "張可麗"
	temp := []rune(name)
	for i := 0;i < len(temp);i++{
		//因為是rune所以要轉型
		fmt.Println(string(temp[i])) //這樣才回印出中文
	}
	```
	
* 上面的範例如果直接用range來取得字串每個值，會自動轉成rune



## 集合們

### Array

#### 練習腳本

arrayPractice

#### 宣告方式

* 基本宣告：[長度]型別

  ``` go
  var myArray [10]int
  ```

  > 長度一定要帶，不然會變切片

* 附帶初始值宣告：[長度]型別{初始值}
  ```go
  var myArray = [10]int{0,1,2,3,4,5,6,7,8,9} //賦予初始值一定要在=後面
  ```

* 利用初始值長度決定陣列長度宣告：[...]型別{初始值}
  ```go
  var myArray = [...]int{1,2,3} //相當於宣告一個長度3的陣列
  ```
  
* 宣告時用索引鍵附值：[長度]型別{索引值：初始值}
	```go
	var myArray = [10]int{0,3:1,6:9}
	```
  
#### 注意事項
* 兩個陣列一定要長度和型別一樣才能比較，不然會直接跳錯
* 其餘陣列用法與TS和C#相同 
* 其餘取值與賦予值得方式一樣



### 切片（slice）

#### 練習腳本

slicePractice

#### 宣告方式

[]型別
```go
var mySlice []int
```

#### 新增元素的方式

* 利用append方式塞值，append會返回一個新的切片

* 新增單一元素：切片 = append(切片,元素)
	```go
	var mySlice []int
	mySlice = append(mySlice,1)
	```
* 一次新增多個元素：切片 = append(切片,元素1,元素2,...)
	```go
	var mySlice []int
	mySlice = append(mySlice,2,6,7)
	```
* 直接把另一個切片塞進一個切片：切片 = append(切片,切片2...)
	```go
	var mySlice []int
	var mySlice2 []int
	
	mySlice = append(mySlice,1,2,3,4)
	mySlice2 = append(mySlice2,mySlice...)
	```
	

#### 注意事項

* 切片只是陣列的包裝，本質還是陣列
* 經過包裝後切片可隨意更改陣列的長度，就像是讓陣列有List的功能一樣
* 因為只是包裝，所以記憶體是指向原本的陣列



## 額外小知識

* strings.Joins可以把切片組合成一組字串，使用方式：strings.Joins(切片,每個元素中間要插入的字元)
* import "os" 可以利用os.Args抓到程式執行時於run後面輸入的訊息，但是會連腳本的路徑一起印出來。
	```go
	//go run . this is test
	func testFunc(){
		fmt.Printf(strings.Joins(os.Args,""))
	}
	```