package main

import (
    "fmt"
    "os"
)

func main() {
	fmt.Printf("Hello, world")
}

// var v1, v2, v3 type
// var v2, v2, v2 type = 1, 2, 3
// var v1, v2, v3 = 1, 2 ,3
// v1, v2, v3 := 1, 2, 3  只能函数内部使用，外部使用var
// _, b = 34, 35 34会直接被丢弃

func ChangeStings() {
	s := "hello"
	c := []byte(s) // 将字符串转换为[]byte类型
	c[0] = "c"
	s2 := string(c)

	s = "c" + s[1:] // 字符串可以切片

	m := "world"
	a := s + m

	b = `hello
              world` // 多行字符串
}

err := errors.New("error")
if err != nil {
    fmt.Print(err)
}

//技巧
//分组
const(
    i = 100
    pi = 3.14
    prefix = "Go_"
)

var(
    i int
    pi float32
    preifx string
)

// iota const默认为0，每行+1， 每遇到一个const关键字，iota就会重置，在同一行值相同
const(
    x = iota // x == 0
    y // y == 1
    z // z == 2
)

// 大写字母开头的变量和函数是可导出的

// 数组, 不能改变长度
// 数组之间的赋值是值的赋值，即把一个数组作为参数传入函数的时候，传入的是该数组的副本，而不是指针
var arr [n]type
a := [10]int{1, 2, 3} // 其他默认为0
c := [...]int{1, 2, 3} //省略长度，go会根据元素来计算
doubleArray := [2][4]int{{1,2,3,4}, {5, 6, 7, 8}}

// slice 切片，动态数组, 少了长度
slice := []byte{'a', 'b', 'c', 'd'}

var ar = [10]byte{'a', 'b', 'c'}
var a, b []byte

a = ar[2:5]
b = ar[3:5] // 数组方括号有数组的长度或... 切片为空
// 切片是引用类型，当引用改变了其中元素的值，其他所有的引用都会改变
// 一个slice像一个结构体，包含一个指针，指向开始的位置，长度，最大长度
// 有len, cap, append, copy 其中append， 如果没有剩余空间，go会动态分配新的空间

slice = array[2:4:7] // 从2-5， 容量是7-2=5

// map 字典, 引用类型， 不是线程安全的， 要使用mutex lock机制
var numbers map[string]int
numbers = make(map[string]int)
numbers["one"] = 1
numbers["two"] = 2

ranting := map[string]float32{"c": 5, "python": 5, "go": 5}
csRanting, ok := ranting["c++"]
if ok {
    fmt.Println(csRanting)
} else {
    pass
}
delete(ranting, "c")

m := make(map[string]string)
m["hello"] = "world"
m1 := m
m1["hello"] = "wcg" // 现在m["hello"] = wcg

// make, new 操作
make 只能用于内建类型(map, slice, channel)的内存分配，会初始化内部的数据结构，填充适当的值，返回的是初始化后的(非零)值
new 用于各种类型的内存分配，返回的是指针

//零值
int int8 int32 int64 rune float32 float64 0
uint byte 0x0
bool false
string ""

if x := computedValue(); x > 10 {
    pass
} else if x < 5 {
    pass
} else {
    pass
}

func goto() {
    i := 0
Here: // 这行的第一个词，以冒号结束作为标签
    println(i)
    i++
    goto Here
}

func sum() {
    sum := 0
    for index := 0; index < 100; index++ {
        sum += index
    }
}

sum := 1
for ; sum < 1000; {
    sum += sum
}

// 可以省略； 相当于while
sum := 1
for sum < 1000 {
    sum += sum
}

for k, v :range map {
    fmt.Println(k)
    fmt.Println(v)
}

i := 10
switch i {
case 1:
    fmt.Pringln(1)
case 2:
    pass
case 10:
    pass
default:
    fmt.Println('default')
}

// 每个case最后带有break, 匹配成功后不会自动执行其他case，而是跳过整个switch，但可以使用fallthrough强制执行后面的代码

// 函数最好返回命名返回值
func SumAndProduct(a, b int) (add, mul int) {
    add = a + b
    mul = a * b
    return
}

// 变参
func myfunc(arg ...int) {
	for _, n := range arg { // arg是一个int的切片
		fmt.Printf(n)
	}
}

// 传值和传指针，传值是值的一个copy，指针才会更改原值
func add1(a int) int {
	a = a + 1
	return a
}
x := 3
x1 := add1(x)
x = 3

func add2(a *int) int { // 仍是copy，只是copy的是一份指针
	*a = *a + 1
	return *a
}

x2 := add2(&x)
x = 4

// 传指针的好处
// 能使多个函数操作同一个对象
// 比较轻量级(8bytes)，只传内存地址，对于大的结构体，用指针可以减少很多的系统开销
// channel, slice, map这三种类型实现机制类似指针，可以直接传递，而不用去地址后传递指针(如改slice的长度，仍需取地址传指针)

//defer 添加多个defer，当函数执行到最后，这些defer会逆序执行， 比如打开文件，自动关闭
func Read() {
	file.Open("flie")
	defer file.Close()
}

// 函数可以作为值和类型

//go中不能抛出异常，而是使用panic和recover机制

// import
import "./model" // 相对路径
import "shorturl/model" // 绝对路径 gopath/src/shroturl/model

import (
    . "fmt" // 可以省略包名 Println("")
)

import (
    f "fmt" // 别名 f.Prinln("")
)

import (
    _ "github.com/..." // 引入该包，不直接使用，而是调用了该包的init函数
)

// 匿名字段
type Human struct {
    name string
    age int
    weight int
}

type Student struct {
    Human
    speciality string
}

makr := Student{Human{"mark", 25, 129}, "hhh"}

// 面向对象，带有接收者的函数，成为method
type Rectagle struct {
    width, height float64
}

type Circle struct {
    radius float64
}

func (r Rectagle) area() float64 {
    return r.width * r.height
}

func (r Circle) area() float64 {
    return
}

// struct只是自定义类型里一种特殊的类型
type ages int
type months map[string]int

// method里的指针，go会自动识别
// method可以继承和重写

func (h *Human) SayHi() {
    fmt.Printf("")
}

func (s *Student) SayHi() {
    pass
}

// interface 是go中设计最精妙的
// interface是一组method签名的组合，通过interface来定义对象的一组行为
// interface定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象实现了此接口
// interface可以被任意对象实现，一个对象可以实现任意多个interface，任何类型都实现了空interface
// go通过interface实现鸭子模型

type Men interface {
    SayHi()
    Sing(lyrics string)
    Guzzle(beerStein string)
}

x := make([]Men, 3)
x[0], x[1], x[2] = mkke, paul, sam

for _, value := range x{
    value.SayHi() // 接口定义的类型变量能存储支持的类型，相当一可以把实现了功能一样的类型可以放在一起。
}

// 一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果返回一个interface{}，那么可以返回任意类型的值

// interface 变量存储的值 value, ok = element.(T)

type Element interface{}
type List [] Element

type Person struct {
    name string
    age int
}

func (p Person) String() string {
    return
}

func inter() {
    list := make(List, 3)
    list[0] = 1
    list[1] = "hello"
    list[2] = Person{"wcg", 25}

    for index, element := range list{
        switch value := element.(type) {
        case int:
            fmt.Printf(int)
        case string:
            fmt.Printf(string)
        case Person:
            fmt.Printf(Person)
        default:
            fmt.Printf("different")
        }
    }
}

//interface真正厉害的是内嵌
//用reflect实现反射，即检查程序在运行时的状态
t := reflect.TypeOf(i) // 获取类型定义里的所有元素
v := reflect.ValueOf(i) // 得到实际的值，还可以去改变

var x float64 = 3.4
p := reflect.ValueOf(&x)
v := p.Elem()
v.SetFloat(7.1)

// 并发
// goroutine, 协程，只占极少的栈内存，4~5k，通过runtime管理
// go 关键字很容易实现并发，上面的多个协程运行在同一个进程中共享数据，不过设计时要使用通道来共享。
// runtime.GOMAXPROCS(n) 设置cpu核数


