package main

import "fmt"

func num(a int) {
	for i := 0; i <= a; i++ {
		fmt.Println(i)
	}
}

func oddprint(a int) {
	for i := 2; i <= a; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func evenprint(a int) {
	for i := 0; i <= a; i++ {
		fmt.Println(i, even(i))
	}
}
