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
