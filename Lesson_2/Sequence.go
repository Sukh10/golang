package main

import "fmt"

func fizz(a int) {
	for i := 0; i <= a; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "buzz")
		} else {
			fmt.Println(i)
		}
	}
}

//0,1,1,2,3.5,8,13
//0+1=1
//1+2 =3
//3+2 = 5
func Fib(num int) {
	a := 0
	b := 1
	fmt.Println(a, b)
	for i := 0; i <= num; i++ {
		sum := a + b     // 0+1 =1
		fmt.Println(sum) //Printed the sum
		a = b            //replace a with b, you don't need a
		b = sum          // then replace b with sum and permanently get rid of a
	}
}

func back(a int) {
	for i := 100; i >= 0; i-- {
		fmt.Println(i)
	}
}

func fives(a int) {
	for i := 100; i >= 5; i = i - 5 {
		fmt.Println(i)
	}
}

func multiples(a int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * a)
	}
}

func star(a int) {
	for i := 0; i <= a; i++ {
		for x := 0; x <= i; x++ {
			fmt.Print(x)
		}
		fmt.Println("")
	}
}

func triangle(a int) {
	for i := 0; i <= a; i++ {
		for x := 0; x <= i; x++ {
			fmt.Print(x)
		}
		fmt.Println("")
	}
	for b := a - 1; b >= 0; b-- {
		for c := 0; c <= b; c++ {
			fmt.Print(c)
		}
		fmt.Println("")
	}
}

func triangle1(a int) {
	size := a * 2
	for i := 0; i <= size; i++ {
		if i <= size/2 {
			for x := 0; x <= i; x++ {
				fmt.Print(x)
			}
		} else {
			for b := 0; b <= a; b++ {
				fmt.Print(b)
			}
			a--
		}
		fmt.Println("")
	}
}

func Prime(a int) {

}
