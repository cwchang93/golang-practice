### 已完成
1. 官網 The tour of go 基本語法 


### 大寫字母開頭為export 
如果要取用package裡面的函數，必須要用大寫
e.g. math.pi => math.Pi

###　function 宣告
- func()
- 參數與結果都要宣告
- 相同型別可以省略

```go

func add(x, y int) int {
	return x + y
}

```

### Print 釐清
1. fmt.Println => print line
    - Print與Println只差在是否有換行

2. fmt.Printf => format的print
    - 第一個參數為string, 


#### :=用途為何?
:=

#### func只要定義就一定會執行?

#### Golang 要點
- 只能用雙引號，不可用單引號
- "%T" type
- 只有var才能 := , const 不能

#### 語法
- number 轉 string
