package main

import "fmt"

func swap(a, b int){
	//first approach
	fmt.Println(a, b)
	a = a + b
	b = a - b
	a = a - b 
	fmt.Println(a, b) 
}

func swap2(a, b int){
	//second approach
	fmt.Println(a, b)
	A := a
	B := b
	A = a
	b = A  
	a = B  
	fmt.Println(a, b)
}

func main(){
	fmt.Println("Hello, Word")
	swap(6, 5)
	swap2(50, 60)
}
