package main // 指明文件属于哪个包， 根据惯例，每个目录只包含一个包

import fm "fmt"

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
