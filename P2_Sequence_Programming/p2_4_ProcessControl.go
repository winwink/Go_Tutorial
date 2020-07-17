package P2_Sequence_Programming


import "fmt"
//2.4 流程控制

// 2.4.1 if, else, else if
func example(x int) int {
	if x==0{
		return 5
	}
	return x
}
- 左花括号必须与if或者else 同一行
- 如果函数有返回值, 在if和else之外, 必须要有一个return

// 2.4.2 switch
switch i{
case 0:
	fmt.Printf("0")
case 1:
	fmt.Printf("1")
case 2:
	fallthrough
case 3:
	fmt.Printf("3")
default:
	fmt.Printf("Default")
}
与C语言相反, C语言需要break打断, 否则就穿透; GO默认打断, 只有加了fallthrough才穿透
switch {
case 0<=Num && Num <=3:
	fmt.Printf("0-3")
case 4<= Num && Num <=6:
	fmt.Printf("4-6")
case 7<= Num && Num<=9:
	fmt.Printf("7-9")
}

// 2.4.3 for
Go只支持for, 不支持while和do-while
sum := 0
for  i:=0;i<10;i++{
	sum += i
}

for{
	sum++
	if sum > 100{
		break
	}
}

// 2.4.4 goto
func myfunc(){
	i :=0
	HERE:
	fmt.Println(i)
	i++
	if i< 10{
		goto HERE
	}
}