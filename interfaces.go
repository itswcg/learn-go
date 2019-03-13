package main

//type Shaper interface {
//	Area() float32
//}
//
//type Square struct {
//	side float32
//}
//
//func (sq *Square) Area() float32 {
//	return sq.side * sq.side
//}
//
//func main() {
//	sq1 := new(Square)
//	sq1.side = 5
//
//	areaInf := sq1
//	fmt.Printf("%f", areaInf.Area())
//}

// 多态
//type Shaper interface {
//	Area() float32
//}
//
//type Square struct {
//	side float32
//}
//
//func (sq *Square) Area() float32 {
//	return sq.side * sq.side
//}
//
//type Rectangle struct {
//	length, width float32
//}
//
//func (r Rectangle) Area() float32 {
//	return r.length * r.width
//}
//
//func main() {
//	r := Rectangle{5, 3}
//	q := &Square{5}
//	shapes := []Shaper{r, q}
//	for n, _ := range shapes {
//		fmt.Println("Shape detail:", shapes[n])
//		fmt.Println("Area of this shape is", shapes[n].Area())
//	}
//}

//type stockPosition struct {
//	ticker     string
//	sharePrice float32
//	count      float32
//}
//
//func (s stockPosition) getValue() float32 {
//	return s.sharePrice * s.count
//}
//
//type car struct {
//	make  string
//	model string
//	price float32
//}
//
//func (c car) getValue() float32 {
//	return c.price
//}
//
//type valuable interface {
//	gerValue() float32
//}
//
//func showValue(asset valuable) {
//	fmt.Printf("Value of the asset is %f\n", asset.gerValue())
//}
//
//func main() {
//	var o valuable = stockPosition{"GooG", 577.20, 4}
//	showValue(o)
//	o = car{"BMW", "M3", 66500}
//	showValue(o)
//}
