package A231001

import "fmt"

func A231001() {
	var a string

	count := 5
	for i := 0; i < count; i++ {
		a = a + "*"
		fmt.Println(a)

	}

}

//
//import "fmt"

/*
func main() {
	fmt.Println("Hello Word")
}



//宣告變數
func main(){
	var a int = 2 //定義 方法一
	//a := 2 //定義 方法二

}


//宣告變數為hello word，並輸出
func main() {
	var a string = "Hello word"

	fmt.Println(a)
}


i++ = i+1


// 直角三角形*，迴圈方式
func main() {
	var a string

	count := 5
	for i := 0; i < count; i++ {

		a = a + "*"
		fmt.Println(a)
		//a = a + "*"

	}
}

//兩者皆相同
b = b+1
b += 1

break 跳脫迴圈

!= 不等於
== 等於
; 分段

v=value
len=length 長度

arr := [5]string{"12","32","42","23","55"}

[0,n)
[包含 )不包含

scanf 接收使用者參數
printf ;F=Format (客製化，需手動換行)
%d、%s、%v  %=告訴空位 d=digital 印數字 s=字串 v=自動判斷
\n 換行符號

&拿地址
rand.Intn() 亂數
scaln(&value)



func main() {

	date := [4]string{"早餐", "午餐", "晚餐", "消夜"}
	name := [4]string{"沙拉", "雞胸肉", "速食", "水果"}

	for i := 0; i < len(date); i++ {
		randFood := rand.Intn(len(name))
		// launch := name[randFood]
		fmt.Printf("%s %s\n", date[i], name[randFood])

	}
}
*/
