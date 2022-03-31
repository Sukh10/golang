package main

import "fmt"

func main() {
	array1(4)
	num := []int{1, 4, 2, 3}
	fmt.Println("The minimum number is : ", min(num))
	fmt.Println("The maximum number is: ", max(num))
	fmt.Println("The sum is: ", sum(num))
	fmt.Println("The average is: ", average(num))
	fmt.Println("Is there a duplicate number: ", duplicate(num))
	fmt.Println("Is there a duplicate number: ", duplicate1(num))
	dup1 := []int{1, 4, 1, 7, 6, 2, 5, 3, 3}
	fmt.Println("The duplicate number is:  ", duplicatenum(dup1))
	fmt.Println("The order is:  ", asc(num))
	fmt.Println("The dsc order is: ", dsc(num))
	test := []int{1, 4, 2, 3, 6}
	fmt.Println("Reverse array is: ", reverse(test))
	test1 := []int{1, 3, 2, -2, -1, 0, -2, 2, 3, 1}
	fmt.Println("Negative numbers: ", negative(test1))

	listOfNames := []string{"sukh", "jag", "mike", "dean", "mike", "preet"}
	fmt.Println("The string is: ", delete1(listOfNames, "mike"))

	fmt.Println("The delete2 is", delete2(listOfNames, "sukh"))

	// listOfNames1 := Person{"sukh", "jag", "mike", "dean", "mike", "preet"}
	// 	fmt.Println("The print:", (listOfNames1("sukh"))

	records := []Person{
		{"sukh"}, {"messi"}, {"ronaldo"}, {"mbappe"},
	}
	fmt.Println("The string is: ", removeobject(records, "messi"))

	findp := []Person1{{"Nick", 18, "Male"}, {"john", 20, "Male"}, {"Priya", 18, "Female"}}
	fmt.Println(findpersonbyname(findp, "john"))
	fmt.Println(findpeoplebyname(findp, 18))

	name := "sukh"
	last := &name
	fmt.Println(*last)

	*last = "sim"
	fmt.Println(name)

	name1 := map[string]string{
		"sukh": "male",
		"ravi": "male",
		"sim":  "female",
		"mani": "male",
	}
	for i, value := range name1 {
		fmt.Println(i, value)
	}

	dup := []int{1, 2, 3, 1, 2, 5, 4}
	unique := nodup(dup)
	fmt.Println(unique)

	listOfPhoneNumbers := []string{"609-122-3921", "800-121-3123", "911-121-1323", "609-231-1231", "911-912-1202", "800-120-1210"}
	GroupByAreacode(listOfPhoneNumbers)

	listOfPeople := []PeopleInfo{{"sukh", 17, "male"}, {"jag", 20, "Male"}, {"ravi", 31, "Male"}, {"Simran", 16, "Female"},
		{"Mani", 30, "Male"}, {"Deep", 50, "Male"}}

	groupByAge(listOfPeople)

	sortAge(listOfPeople)

}
