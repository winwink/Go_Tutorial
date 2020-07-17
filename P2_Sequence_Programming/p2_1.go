package P2_Sequence_Programming

// 0 
// 小写字母表示private, 大写字母表示public

// 2.1.1 变量声明
var v1 int
var v2 string
var v3 [10]int //数组
var v4 []int //数组切片
var v5 struct {
	f int
}
var v6 *int //指针
var v7 map[string]int
var v8 func(a int) int
var {
	v9 int
	v10 string
}

// 2.1.2 变量初始化
var v11 int = 10
var v12 = 10
v13 := 10

// 2.1.3 变量赋值
var v14 int
v14 = 123
i, j = j, i

// 2.1.4 匿名变量
func GetName() (firstName, lastName, nickName string){
	return "May", "Chan", "Chibi Maruko"
}

_, _, nickName := GetName()

// 2.2.1 字面常量
-12
3.1415926 //浮点常量
3.2+12i //复数常量
true //布尔常量
"foo" //字符串常量

// 2.2.2 常量定义
const Pi float64 = 3.1415926 
const zero = 0.0 //无类型浮点常数
const {
	size int64 = 1024
	eof = -1 //无类型整型常数
}

const u, v float32 = 0, 3 // u=0.0, v=3.0
const a, b, c = 3, 4, "foo" // a=3, b=4, c="foo" 无类型整型和字符串常量

// 2.2.3 预定义常量
true, false, iota
const {
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
}
const {//const省略赋值, 与上面的表达式相同
	c0 = iota // c0 == 0
	c1				// c1 == 1
	c2				// c2 == 2 
}

// 2.2.4 枚举
const {
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays // 包内可见, 包外无法访问
}

// 2.3类型
基础类型
bool
int8, byte, int16, int, uint, uintptr //一般用int和uint
float32, float64
complex64, complex128
string
rune //char
error
复合类型
pointer // var v6 *int
array // var v7 [6]int
slice // var v8 []int
map
chan
struct
interface

// 2.3.1 bool
var v1 bool
v1 = true
v2 := (1 == 2)

// 2.3.2 int
var int value1

// 2.3.3 float32 float64
var fvalue1
fvalue1 := 1.0

// 2.3.4 complex
value2 := 3.2 + 12i
value3 := complex(3.2, 12)

// 2.3.5 string
var str string
str = "Hello World"
ch := str[0] // get first char of string

str := "Hello World"
str[0] = 'X' // complie error

x + y //string combine
len(s) // get length of string
s[i]  //get char of string

str := "Hello, 世界"
n := len(str) //n=13, chinese char contains 3 utf-8 code
for i:=0; i<n;i++{
	ch := str[i]
	fmt.Println(i, ch) //print ascii code of char
}
for i, ch := range str{  // i=9
	fmt.Println(i, ch) // ch's type is rune, chinese char contains 3 utf-8 code
}

// 2.3.6 char & rune
byte 代表utf-8字符串的单个字节的值
rune 代表单个Unicode字符

// 2.3.7 array
[32]byte //长度32的数组, 每个元素为一个byte
[2*N] struct {x,y int32} //复杂类型数组
[1000]*float64 //指针数组
[3][5]int //二维数组

for i:=0; i< len(array);i++{
	fmt.Println("Element", i, "of array is", array[i])
}

for i, v := range array{
	fmt.Println("Element", i, "of array is", v)
}

数组是值类型, 传递后只是复制, 无法进行修改

// 2.3.7 slice
基于数组创建切片
var myArray [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
var mySlice []int = myArray[:5]//前5个
for _,v := range mySlice {
	fmt.print(v, " ") //1 2 3 4 5
}

直接创建
mySlice1 := make([]int, 5)//初始元素个数为5, 初始值为0
mySlice2 := make([]int, 5, 10)//初始元素个数为5, 初始值为0, capacity为10, 可以通过cap(mySlice)查看容量
mySlice3 := []int{1,2,3,4,5}

遍历
for i:=0;i<len(mySlice);i++{
	fmt.Println("mySlice[",i,"] = ",mySlice[i])
}
for i, v:=range mySlice{
	fmt.Println("mySlice[",i,"] = ",v)
}

mySlice = append(mySlice, 1, 2, 3)
mySliec = append(mySlice, mySlice1)

slice1 := []int{1,2,3,4,5}
slice2 := []int{6,7,8}
copy(slice2, slice1)//copy(destination, source) slice2: {1,2,3}
copy(slice1, slice2)// slice1:{6,7,8,4,5}

//2.3.9 map
type PersonInfo struct {
	ID string
	Name string
	Address string
}

var personDB map[string] PersonInfo
personDB = make(map[string] PersonInfo)

personDB["12345"] = PersonInfo("12345", "Tom" ,"Room 203")
personDB["1"] = PersonInfo("1", "Jack", "Room 101")

person, ok := personDB["1234"]

if ok{
	fmt.Println("Found People ",person.Name," with ID 1234")
} else {
	fmt.Println("Did not find people with ID 1234")
}

delete(personDB, "1234")
