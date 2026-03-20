package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type (
	By           func(p1, p2 *Person) bool
	PersonSorter struct {
		People []Person
		By     func(p1, p2 *Person) bool
	}
)

func (s *PersonSorter) Len() int {
	return len(s.People)
}

func (s *PersonSorter) Less(i, j int) bool {
	return s.By(&s.People[i], &s.People[j])
}

func (s *PersonSorter) Swap(i, j int) {
	s.People[i], s.People[j] = s.People[j], s.People[i]
}

func (by By) Sort(people []Person) {
	ps := &PersonSorter{
		People: people,
		By:     by,
	}

	sort.Sort(ps)
}

// type (
// 	By     func(p1, p2 *Person) bool
// 	ByAge  []Person
// 	ByName []Person
// )

// func (a ByAge) Len() int {
// 	return len(a)
// }

// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }

// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

// func (a ByName) Len() int {
// 	return len(a)
// }

// func (a ByName) Less(i, j int) bool {
// 	return a[i].Name < a[j].Name
// }

// func (a ByName) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

func main() {
	// numbers := []int{5, 4, 3, 2, 1}
	// sort.Ints(numbers)
	// fmt.Println("Sorted Numbers:", numbers)

	// stringSlices := []string{"John", "Anthony", "Steve", "Victor", "Walter"}
	// sort.Strings(stringSlices)
	// fmt.Println("Sorted Strings:", stringSlices)

	people := []Person{
		{
			"Alice",
			30,
		},
		{
			"Bob",
			25,
		},
		{
			"Anna",
			35,
		},
	}

	fmt.Println("Unsorted by Age:", people)
	age := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	By(age).Sort(people)
	fmt.Println("Sorted by Age: ", people)

	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	By(name).Sort(people)
	fmt.Println("Sorted by Name:", people)

	ageDesc := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}
	By(ageDesc).Sort(people)
	fmt.Println("Sorted by Age [DESC]:", people)

	lenName := func(p1, p2 *Person) bool {
		return len(p1.Name) < len(p2.Name)
	}

	By(lenName).Sort(people)
	fmt.Println("Sorted by Length of Name:", people)

	// ============== SORT.Slice ==============
	stringSlice := []string{"banana", "apple", "cherry", "grapes", "guava"}
	sort.Slice(stringSlice, func(i, j int) bool {
		// Sort by last character
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})
	fmt.Println("Sorted by Last Character:", stringSlice)
}
