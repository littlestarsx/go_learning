/**
Go 指南
 */
package main

//func main() {
//	fmt.Println("Hello World");
//}


//包
//func main() {
//	fmt.Println("My favorite number is,", rand.Intn(10))
//}


//导入
//func main() {
//	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))//开平方
//}


//导出名
//func main() {
//	fmt.Println(math.Pi)
//}

//函数
//func add (x int, y int) int {
//	return x + y
//}
//
//func main() {
//	fmt.Println(add(12,  30))
//}


//函数 续 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略
//func add (x, y int) int {
//	return x + y
//}
//
//func main() {
//	fmt.Println(add(12,  30))
//}


//多值返回
//func swap(x, y string) (string, string) {
//	return x, y
//}
//
//func main() {
//	//a, b := swap("hello", "world")
//	var a, b string
//	a, b = swap("hello", "world")
//	fmt.Println(a, b)
//}


//命名返回值
//func split(sum int) (x, y int){
//	x = sum * 4 / 9
//	y = sum - x
//	return
//}

//func split(sum int) (x, y int){
//	x = sum * 4 / 9
//	y = sum - x
//	return x, y
//}
//
//func main() {
//	fmt.Println(split(3))
//}


//变量
//var c, python, java bool
//
//func main() {
//	var i int
//	fmt.Println(i, c, python, java)
//}


//变量初始化
//var i, j int = 1, 2
//
//func main() {
//	var c, python, java = true, false, "no!"
//	fmt.Println(i, j, c, python, java)
//}

//短变量声明
//func main() {
//	var i, j int = 1, 2
//	k := 3
//	c, python, java := true, false, "no!"
//	fmt.Println(i, j, k, c, python, java)
//}


//基本类型
//var (
//	ToBe bool = false
//	MaxInt uint64 = 1<<64 - 1
//	z complex128 = cmplx.Sqrt(-5 + 12i)
//)
//
//func main() {
//	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
//	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
//	fmt.Printf("Type: %T Value: %v\n", z, z)
//}


//零值 没有明确初始值的变量声明会被赋予它们的 零值。
//数值类型为 0，
//布尔类型为 false，
//字符串为 ""（空字符串）。
//func main() {
//	var i int
//	var f float64
//	var b bool
//	var s string
//	fmt.Printf("%v %v %v %q\n", i, f, b, s)
//}


//类型转换
//func main() {
//	var x, y int = 3, 4
//	var f float64 = math.Sqrt(float64(x*x + y*y))
//	var z uint = uint(f)
//	fmt.Println(x, y, z)
//}


//类型推导
//func main() {
//	v := 42 // 修改这里！
//	fmt.Printf("v is of type %T\n", v)
//}


//常量
//const Pi = 3.14
//
//func main() {
//	const World = "世界"
//	fmt.Println("Hello", World)
//	fmt.Println("Happy", Pi, "Day")
//
//	const Truth = true
//	fmt.Println("Go rules?", Truth)
//}


//数值常量
//const (
//	Big = 1 << 100
//	Small = Big >> 99
//)
//
//func needInt(x int) int {
//	return x * 10 + 1
//}
//func needFloat(x float64) float64 {
//	return x * 0.1
//}
//
//func main() {
//	fmt.Println(Small)
//	var BigValue float64 = Big
//	fmt.Println(BigValue)
//	fmt.Println(needInt(Small))
//	fmt.Println(needFloat(Small))
//	fmt.Println(needFloat(Big))
//}