package main

import "fmt"

// func delete1(listOfNames []string, name string) []string {
// 	remove := []string{}
// 	for _, a := range listOfNames {
// 		if a != name {
// 			remove = append(remove, a)
// 		}
// 	}
// 	return remove
// }
func delete1(listOfNames []string, name string) []string {
	for i, a := range listOfNames {
		if a == name {

			listOfNames = append(listOfNames[:i], listOfNames[i+1:]...)
		}
	}
	return listOfNames
}

func delete2(listOfNames []string, name string) []string {
	p1 := []string{}
	for _, a := range listOfNames {
		if a != name {
			p1 = append(p1, a)
		}
	}
	return p1
}

type Person struct {
	Name string
}

func (b Person) Getname(name string) string {
	return b.Name
}

func removeobject(records []Person, name string) []Person {
	p1 := []Person{}
	for _, p := range records {
		if p.Name != name {
			p1 = append(p1, p)
		}
	}
	return p1
}

type Person1 struct {
	Name   string
	Age    int
	Gender string
}

func findpersonbyname(findp []Person1, name string) *Person1 {
	for _, b := range findp {
		if b.Name == name {
			return &b
		}
	}
	return nil
}

func findpeoplebyname(findp []Person1, age int) []Person1 {
	sameAge := []Person1{}
	for _, b := range findp {
		if b.Age == age {
			sameAge = append(sameAge, b)
		}
	}
	return sameAge
}

func nodup(dup []int) []int { //{1, 2, 3, 1, 2, 5, 4}
	d := []int{}
	nod := map[int]int{}
	for a := range dup {
		//fmt.Println(nod[dup[a]])
		//if _, found := nod[dup[a]]; found {
		//	continue
		//}
		//nod[dup[a]] = 0
		//d = append(d, dup[a])

		if nod[dup[a]] != 1 {
			fmt.Println(nod[dup[a]])
			nod[dup[a]] = 1
			d = append(d, dup[a])
		}
	}
	return d
}

func GroupByAreacode(listOfPhoneNumbers []string) {
	//p := []string{}
	cell := make(map[string][]string, 1)
	for _, phoneNumber := range listOfPhoneNumbers {
		areacode := phoneNumber[0:3]
		if phoneNumbers, found := cell[areacode]; found { // if phomemumber with that area code is found, then add all numbers starting with that areacode
			phoneNumbers = append(phoneNumbers, phoneNumber)
			cell[areacode] = phoneNumbers
			continue
		}
		var phoneNumbers []string
		phoneNumbers = append(phoneNumbers, phoneNumber)
		cell[areacode] = phoneNumbers
	}
	fmt.Println(cell)
}

type PeopleInfo struct {
	Name   string
	Age    int
	Gender string
}

func groupByAge(listOfPeople []PeopleInfo) {
	class := make(map[string][]PeopleInfo, 1)
	info := []PeopleInfo{}
	info1 := []PeopleInfo{}
	info2 := []PeopleInfo{}
	minor := "minor"
	//var min int
	for _, person := range listOfPeople {
		if person.Age < 18 {
			info = append(info, person)
			class[minor] = info
			continue
		}
		if person.Age > 18 && person.Age < 40 {
			info1 = append(info1, person)
			class["adult"] = info1
		}
		if person.Age > 40 {
			info2 = append(info2, person)
			class["seniorcitizen"] = info2
		}
	}
	fmt.Println(class)
}

func sortAge(listOfPeople []PeopleInfo) {
	people := []int{}
	for _, person := range listOfPeople {
		people = append(people, person.Age)
	}
	for i := 0; i < len(people); i++ {
		for b := i + 1; b < len(people); b++ {
			if people[b] < people[i] {
				min := people[b]
				people[b] = people[i]
				people[i] = min
			}
		}
	}
	fmt.Println(people)
}
