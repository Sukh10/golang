package main

import (
	"fmt"
)

func array1(a int) {
	num := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < a; i++ {
		fmt.Println(num[i])
	}
}

func min(num []int) int {
	min := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] < min {
			min = num[i]
		}
	}
	return min
}
func max(num []int) int { //{1, 4, 2, 3}
	max := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] > max {
			max = num[i]
		}
	}
	return max
}

func sum(num []int) int {
	add := 0
	for i := 0; i <= len(num); i++ {
		add += i
	}
	return add
}

func average(num []int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += i
	}
	avg := sum / len(num)
	return avg
}

func duplicate(num []int) bool {
	duplicate := num[0]
	for i := 1; i < len(num); i++ {
		if duplicate == num[i] {
			return true
		}
		for b := i + 1; b < len(num); b++ {
			if num[i] == num[b] {
				return true
			}
		}
	}
	return false
}

func duplicate1(num []int) bool {
	for i := 0; i < len(num)-1; i++ {
		for b := i + 1; b < len(num); b++ {
			if num[i] == num[b] {
				return true
			}
		}
	}
	return false
}

func duplicatenum(num []int) []int {
	dupnum := []int{}
	for i := 0; i < len(num); i++ {
		for b := i + 1; b < len(num); b++ {
			if num[i] == num[b] {
				dupnum = append(dupnum, num[i])
			}
		}
	}
	return dupnum
}

func asc(num []int) []int {
	for i := 0; i < len(num); i++ { //{3, 2, 1, 4}
		for b := i + 1; b < len(num); b++ {
			if num[b] < num[i] {
				min := num[b] //{2,3}
				num[b] = num[i]
				num[i] = min
			}
		}
	}
	return num
}

func dsc(num []int) []int {
	for i := 0; i < len(num); i++ { //{3, 2, 1, 4}
		for b := i + 1; b < len(num); b++ {
			if num[b] > num[i] {
				min := num[b] //{2,3}
				num[b] = num[i]
				num[i] = min
			}
		}
	}
	return num
}

//{1, 4, 2, 3}
func reverse(num []int) []int {
	for i := 0; i <= len(num)/2-1; i++ {
		temp := num[i]
		num[i] = num[len(num)-1-i]
		num[len(num)-1-i] = temp
	}
	return num
}

func negative(num []int) []int {
	positive := []int{}
	negative := []int{}
	for i := 0; i < len(num); i++ {
		if num[i] > 0 {
			positive = append(positive, num[i])
		} else if num[i] < 0 {
			negative = append(negative, num[i])
		}
	}
	return append(negative, positive...)
}
