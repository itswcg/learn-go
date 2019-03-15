package main // 指明文件属于哪个包， 根据惯例，每个目录只包含一个包

import (
	fm "fmt"
)

type (
	IZ  int
	FZ  float64
	STR string
)

func main() {
	fm.Println("Hello, world")
}

//var a iZ = 5

func Sum(a, b int) int {
	return a + b
}

// go的文件命名均由小写字母组成
// 类型 int float bool string 结构化的： struct, array, slice, map, channel, interface
// 结构化的类型没有真正的值，可以用nil作为默认值
// func FunctionName(a typea, b typeb) (t1 type1, t2 type2)
// 类型转换 a := 5.0 b := int(a)
// const PI = 3.14

//const (
//	a = iota iota=0, 新的一行使用时，自动加1
//	b
//	c
//)

//var a int 变量命名使用驼峰, 如果希望被外部包使用，那单词首字母也大写
//var (
//	a int
//	b bool
//	str sting
//)

//var a = 15  自动推断其类型， 全局
//a := 1 声明局部变量

// int, float, bool, string属于值类型， 指针, slices, maps, channel属于引用类型
// 1e3 = 1000 e表示10的连乘
// 格式化 %d 整数 %x 16进制表示的数字 %g 浮点数 %f 输出浮点数 %e 科学计数表示法
// %0d 规定输出定长的整数 %n.mg 表示数字n并精确到小数点m位 %v 复数 %b 表示位
// %c 输出字符 %U 输出U+hhhh的字符串

//var cl complex64 = 5 + 10i 复数
//字符只是整数的特殊用例,用双引号, \n 换行 \r 回车 \t tab键 \u Unicode字符 \\ 反斜杠自身
// `` 非解释字符串

// 字符串
//strings.HasPrefix(s, prefix string)
//strings.HasSuffix(s, suffix string)
//strings.Contains(s, substr, string)
//strings.Index(s, str string)
//strings.LastIndex(s, str string)
//strings.IndexRune(s string, ch int)
//strings.Replace(str, old, new, n)
//strings.Count(s, str string)
//strings.Repeat(s, count int)
//strings.ToUpper(s)
//strings.TrimSpace(s)
//strings.Split(s, sep)
//Strings.Join(sl []string, sep string)
//Strings.NewReader(str)
//Strings.Read()
//strconv.Itoa(i int)
//strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int)
//strconv.Atoi(s string)
//strconv.ParseFloat(s string, bitSize int) (f float64, err error)

//指针， 能控制指针，但不能指针运算
//var intP *int
//intp = &il %p 指正格式化
// 使用一个指针引用一个值时称为间接引用

//函数
//go不支持函数重载的原因是需要多余的类型匹配影响性能
//按值传递，传递的是副本，不会影响原来的值， 按引用传递，&variable, 指针的值(一个地址)会被复制，但指针的值
//所指向的地址上的值不会被复制，我们可以通过修改指针的值来修改这个值所指向的地址的值
//空白符 _ 丢弃值
//传递指针不但可以节省内存，而且赋予函数直接修改外部变量的能力

//func Min(a ...int) int {} 变参函数
//arr := []int{1, 2, 3}
//x = Min(arr...)

//defer 延缓 允许我们进行一些函数执行完后的收尾工作
// 匿名函数 fplus := func(x, y int) int {return x + y} fplus(3, 4)
// 工厂函数 一个返回值为另一个函数的函数

//数组 最长2GB var arr1 [5]int
//for i,_ := range arr1 {}
//var arrAge = [5]int{1, 2, 3}
//var arrLazy = [...]int{5, 6, 7, 8}
//var arr = [5]string{3: "chris", 4: "ron"}
//var arr = [3][5]int
//将一个大数组传递给函数会消耗很多内存， 1：传递数组的指针 2：使用数组的切片
//切片相当于list，是一个长度可变的数组
//var slice1 []type = arr1[:] slice1 = &arr1
//s = s[:cap(s)] 扩展到它的大小上限
//slice1 := make([]type, len)
// new(T)为每个新的类型T分配一个内存，返回指针，适用于值类型
// make(T)返回T的初始值，适用于引用类型：切片，map，channel

//bytes var buffer bytes.Buffer var r *bytes.Buffer = new(bytes.Buffer)
//slice1 := make([]type, start_length, capacity) 切片的初始长度， 容量
//扩容 sl = sl[0:len(sl)+1]
//x = append(x, y)
//map = python中的字典
//val1, isPresent = map[key1] 判断是否存在
//获取一个map的切片，要使用两次make

//结构体(ADT 抽象数据类型)
//type identifier struct {
//	field1 type1
//	field2 type2
//}
//var t *T = new(T)
//t := new(T)
//type myStruct struct { i int}
//var v myStruct v是结构体类型变量
//var p *myStruct p是指向一个结构体类型变量的指针
//ms := &struct1{10, 15.5, "wcg"} 或
//var mt struct1  ms := struct1{10, 15.5, "wcg"}
// new(Type) 和 &Type{} 是等价的
//type Node struct { 链表
//	data float64
//	su   *Node
//}
//type Node struct { 双向链表
//	pr   *Node
//	data float64
//	su   *Node
//}
//type Tree struct { 二叉树
//	le   *Tree
//	data float64
//	ri   *Tree
//}
//结构体工厂，它返回一个指向结构体实例的指针，实例了类型的一个对象
//type File struct {
//	fd   int
//	name string
//}
//func NewFile(fd int, name string) *File {
//	if fd < 0 {
//		return nil
//	}
//
//	return &File(fd, name)
//}
//f := NewFile(10, "./test.txt")
// make() slices/maps/channels new() struct

//go语言的继承是通过内嵌和组合实现的，组合更受青睐
//在一个结构体中对于每一种数据类型只能有一个匿名字段

//函数将变量作为参数，方法在变量上被调用

//在go中，类型就是类，go没有类继承的概念
//继承有两个好处：代码复用和多态
//go中，代码复用通过组合和委托实现，多态通过接口的使用来实现，称为组件编程

// go 接口： 谁能搞定这件事，它就可以用在这儿
// 类型断言
//if v, ok := varl.(T); ok {}

//运行时的异常 panic("A server error occurred")
//if err != nil {
//	panic("")
//}
