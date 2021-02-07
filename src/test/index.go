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


//for循环
//func main() {
//	sum := 0
//	for i := 0; i < 10; i++ {
//		sum += i
//	}
//	fmt.Println(sum)
//}


//for（续）初始化语句和后置语句是可选的
//func main() {
//	sum := 1
//	for ; sum < 1000; {
//		fmt.Println(sum)
//		sum += sum
//	}
//	fmt.Println(sum)
//}


//for 是 Go 中的 “while”
//此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。
//func main() {
//	sum := 1
//	for sum < 1000 {
//		fmt.Println(sum)
//		sum += sum
//	}
//	fmt.Println(sum)
//}


//无限循环
//func main() {
//	for {
//		fmt.Println()
//	}
//}


//if Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。
//func sqrt(x float64) string {
//	if x < 0 {
//		return sqrt(-x) + "i"
//	}
//	return fmt.Sprint(math.Sqrt(x))
//}
//
//func main () {
//	fmt.Println(sqrt(2), sqrt(-4))
//}


//if的简短语句
//func pow(x, n, lim float64) float64 {
//	if v :=math.Pow(x, n); v < lim {//math.Pow 求几次方
//		return v
//	}
//	return lim
//}
//
//func main() {
//	fmt.Println(pow(3, 2, 10), pow(3,3, 20))
//}


//if 和 else
//func pow(x, n, lim float64) float64 {
//	if v :=math.Pow(x, n); v < lim {//math.Pow 求几次方
//		return v
//	} else {
//		fmt.Println("%g > %g", v, lim)
//	}
//	// 这里开始就不能使用 v 了
//	return lim
//}
//
//func main() {
//	fmt.Println(pow(3, 2, 10), pow(3,3, 20))
//}


//switch
//func main() {
//	fmt.Print("Go runs on ")
//	switch os := runtime.GOOS; os {
//	case "darwin":
//		fmt.Println("OS X.")
//	case "linux":
//		fmt.Println("Linux.")
//	default:
//		// freebsd, openbsd,
//		// plan9, windows...
//		fmt.Printf("%s.\n", os)
//	}
//}


//switch 的求值顺序
//func main() {
//	fmt.Println("When's Saturday?")
//	today := time.Now().Weekday()
//	fmt.Println(today)
//	switch time.Saturday {
//	case today + 0:
//		fmt.Println("Today.")
//	case today + 1:
//		fmt.Println("Tomorrow.")
//	case today + 2:
//		fmt.Println("In two days.")
//	default:
//		fmt.Println("Too far away.")
//	}
//}


//没有条件的 switch
//func main() {
//	t := time.Now()
//	fmt.Println(t.Hour())
//	switch {
//	case t.Hour() < 12:
//		fmt.Println("Good morning!")
//	case t.Hour() < 17:
//		fmt.Println("Good afternoon.")
//	default:
//		fmt.Println("Good evening.")
//	}
//}


//defer defer
//语句会将函数推迟到外层函数返回之后执行。
//推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
//func main() {
//	defer fmt.Println("world")
//
//	fmt.Println("hello")
//}


//defer 栈
//推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
//func main() {
//	fmt.Println("counting")
//
//	for i := 0; i < 10; i++ {
//		defer fmt.Println(i)
//	}
//
//	fmt.Println("done")
//}


//