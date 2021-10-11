package main

import "fmt"

func sum(numbers ...int) (total int) {
	for _, n := range numbers {
		total += n
	}
	return
}

func ex1() {
	nums := []int{1, 2, 3, 4, 5}
	summed := sum(nums...)
	fmt.Println(summed)
}

func evenOdd(num int) (int, bool) {
	halved := num / 2
	if halved%2 == 0 {
		return halved, false
	} else {
		return halved, true
	}
}

func ex2() {
	fmt.Println(evenOdd(1))
	fmt.Println(evenOdd(2))
}

func makeOddGenerator() (gen func() int) {
	i := 1
	gen = func() (ret int) {
		ret = i
		i += 2
		return
	}
	return
}

func ex3() {
	oddGen := makeOddGenerator()
	fmt.Println(oddGen())
	fmt.Println(oddGen())
	fmt.Println(oddGen())
}

func square(x *float64) {
	*x = *x * *x
}

func ex4() {
	x := 1.5
	square(&x)
	fmt.Println(x)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func ex5() {
	fmt.Println("fib")
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
}

func swap(x, y *int) {
	_x := *x
	_y := *y
	*x = _y
	*y = _x
}

func ex11() {
	x, y := 1, 2
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex11()
}
